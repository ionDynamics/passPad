package persistence

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"go.iondynamics.net/iDhelper/crypto"
	"go.iondynamics.net/iDlogger"

	"go.iondynamics.net/passPad/account"
	"go.iondynamics.net/passPad/vault"
)

var db *bolt.DB

func Init(boltPath string) {
	var err error
	db, err = bolt.Open(boltPath, 0600, nil)
	if err != nil {
		iDlogger.Panic(err)
	}
}

func Close() {
	db.Close()
}

func GetVault(id, secret string) (vault.Vault, error) {
	var v vault.Vault
	err := db.View(func(tx *bolt.Tx) error {
		vaults := tx.Bucket([]byte("vaults"))
		if vaults == nil {
			return errors.New("no vaults bucket")
		}
		byt := vaults.Get([]byte(id))
		if len(byt) < 1 {
			return errors.New("no such vault")
		}
		return json.Unmarshal([]byte(crypto.Decrypt(secret, string(byt))), &v)
	})
	return v, err
}

func SetVault(id, secret string, v vault.Vault) error {
	return db.Update(func(tx *bolt.Tx) error {

		byt, err := json.Marshal(v)
		if err != nil {
			return err
		}

		vaults, err := tx.CreateBucketIfNotExists([]byte("vaults"))
		if err != nil {
			return err
		}

		if secret == "" {
			panic("empty secret")
		}
		return vaults.Put([]byte(id), []byte(crypto.Encrypt(secret, string(byt))))
	})
}

func SetAccount(id string, a *account.Account) error {
	return db.Update(func(tx *bolt.Tx) error {

		byt, err := json.Marshal(a)
		if err != nil {
			return err
		}

		accounts, err := tx.CreateBucketIfNotExists([]byte("accounts"))
		if err != nil {
			return err
		}

		return accounts.Put([]byte(id), []byte(crypto.Encrypt(a.Pass, string(byt))))
	})
}

func GetAccount(id, pass string) (*account.Account, error) {
	a := &account.Account{Vaults: make(map[string]string)}
	err := db.View(func(tx *bolt.Tx) error {
		accounts := tx.Bucket([]byte("accounts"))
		if accounts == nil {
			return errors.New("no accounts bucket")
		}
		byt := accounts.Get([]byte(id))
		if len(byt) < 1 {
			return errors.New("no such account")
		}
		return json.Unmarshal([]byte(crypto.Decrypt(pass, string(byt))), a)
	})
	if err == nil {
		if a.Vaults == nil {
			a.Vaults = make(map[string]string)
		}
	}
	return a, err
}

func GetPublicKey(id string) (string, error) {

	var publicKey string

	err := db.View(func(tx *bolt.Tx) error {
		publicKeys := tx.Bucket([]byte("public_keys"))
		if publicKeys == nil {
			return errors.New("no public keys bucket")
		}
		byt := publicKeys.Get([]byte(id))
		if len(byt) < 1 {
			return errors.New("no such public key")
		}
		publicKey = string(byt)
		return nil
	})

	return publicKey, err
}

func SetPublicKey(id string, publicKey string) error {
	return db.Update(func(tx *bolt.Tx) error {

		publicKeys, err := tx.CreateBucketIfNotExists([]byte("public_keys"))
		if err != nil {
			return err
		}

		return publicKeys.Put([]byte(id), []byte(publicKey))
	})
}

func AccountExists(name string) bool {
	ret := false
	db.View(func(tx *bolt.Tx) error {
		accounts := tx.Bucket([]byte("accounts"))
		if accounts == nil {
			return errors.New("no accounts bucket")
		}
		byt := accounts.Get([]byte(name))
		if len(byt) < 1 {
			ret = false
		} else {
			ret = true
		}
		return nil
	})
	return ret
}

func VaultExists(identifier string) bool {
	ret := false
	db.View(func(tx *bolt.Tx) error {
		vaults := tx.Bucket([]byte("vaults"))
		if vaults == nil {
			return errors.New("no vaults bucket")
		}
		byt := vaults.Get([]byte(identifier))
		if len(byt) < 1 {
			ret = false
		} else {
			ret = true
		}
		return nil
	})
	return ret
}
