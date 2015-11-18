package passpad

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/pquerna/otp/totp"
	"go.iondynamics.net/iDhelper/randGen"
	"go.iondynamics.net/passPad/account"
	"go.iondynamics.net/passPad/config"
	"go.iondynamics.net/passPad/persistence"
	"go.iondynamics.net/passPad/vault"
	"image/png"
)

func AuthAccount(u, p string) *account.Account {
	a, err := persistence.GetAccount(u, p)
	if err != nil {
		return nil
	} else {

		if len(a.PrivateKey) < 1 {
			pk, err := rsa.GenerateKey(rand.Reader, 2048)

			if err == nil {

				a.PrivateKey = pem.EncodeToMemory(&pem.Block{
					Type:  "RSA PRIVATE KEY",
					Bytes: x509.MarshalPKCS1PrivateKey(pk),
				})

				persistence.SetAccount(a.User, a)

				pubASN1, _ := x509.MarshalPKIXPublicKey(&pk.PublicKey)

				publicKey := pem.EncodeToMemory(&pem.Block{
					Type:  "RSA PUBLIC KEY",
					Bytes: pubASN1,
				})

				persistence.SetPublicKey(a.User, string(publicKey))
			}
		}

		return a
	}
}

func ValidToken(acc *account.Account, token string) bool {
	return totp.Validate(token, acc.Secret)
}

func RegisterAccount(u, p string) error {
	if AccountExists(u) {
		return errors.New("account already exists")
	}

	a := &account.Account{User: u, Pass: p}
	return persistence.SetAccount(a.User, a)
}

func AccountExists(name string) bool {
	return persistence.AccountExists(name)
}

func AccountSetup(acc *account.Account) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      config.Std.PassPad.OtpIssuer,
		AccountName: acc.User,
	})
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	img, err := key.Image(250, 250)
	if err != nil {
		return "", err
	}
	png.Encode(&buf, img)
	b := base64.StdEncoding.EncodeToString(buf.Bytes())
	acc.Secret = key.Secret()
	acc.ValidSecret = false
	err = persistence.SetAccount(acc.User, acc)
	return b, err
}

func ValidateAccount(acc *account.Account, token string) error {
	if ValidToken(acc, token) {
		acc.ValidSecret = true
		return persistence.SetAccount(acc.User, acc)
	} else {
		return errors.New("invalid token")
	}
}

func CreateVault(acc *account.Account, title string, description string) (vault.Vault, string, error) {
	v := vault.New()
	v.Access = append(v.Access, acc.User)
	v.Title = title
	v.Description = description
	secret := randGen.String(256)
	acc.Vaults[v.Identifier] = secret
	err := persistence.SetVault(v.Identifier, acc.Vaults[v.Identifier], v)
	if err != nil {
		return v, secret, err
	}
	err = persistence.SetAccount(acc.User, acc)
	return v, secret, err
}

func ListVaults(acc *account.Account) (vaults []vault.Vault, err error) {
	for identifier, secret := range acc.Vaults {
		var v vault.Vault
		v, err = persistence.GetVault(identifier, secret)
		if err != nil {
			return
		}
		vaults = append(vaults, v)
	}
	return
}

func OpenVault(acc *account.Account, identifier string) (vault.Vault, error) {
	secret, ok := acc.Vaults[identifier]
	if !ok {
		return vault.Vault{}, errors.New("no access")
	}
	return persistence.GetVault(identifier, secret)
}

func UpsertVault(acc *account.Account, identifier, title string, description string) error {
	var v vault.Vault
	var err error
	var secret string

	if identifier == "" {
		v, secret, err = CreateVault(acc, title, description)
	} else {
		v, err = OpenVault(acc, identifier)
	}

	if err != nil {
		return err
	}

	v.Description = description

	return persistence.SetVault(v.Identifier, secret, v)
}

func UpsertEntry(acc *account.Account, identifier, name, user, pass, url string, data map[string]string) error {
	secret, ok := acc.Vaults[identifier]
	if !ok {
		return errors.New("no access")
	}

	v, err := persistence.GetVault(identifier, secret)
	if err != nil {
		return err
	}

	v.Entries[name] = vault.Entry{
		Name: name,
		User: user,
		Pass: pass,
		Url:  url,
		Data: data,
	}

	return persistence.SetVault(identifier, secret, v)
}
