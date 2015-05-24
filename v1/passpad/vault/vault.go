package vault

import (
	"go.iondynamics.net/iDhelper/randGen"
)

type Entry struct {
	Name string
	User string
	Pass string
	Url  string
	Data map[string]string
}

type Vault struct {
	Identifier  string
	Description string
	Access      []string
	Entries     map[string]Entry
}

func New() Vault {
	return Vault{
		Identifier: randGen.String(64),
		Entries:    make(map[string]Entry),
	}
}
