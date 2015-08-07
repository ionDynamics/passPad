package account

type Account struct {
	User          string
	Pass          string
	Secret        string
	ValidSecret   bool
	RecoveryCodes []string
	Vaults        map[string]string
	PrivateKey    string
}
