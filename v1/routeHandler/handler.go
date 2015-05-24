package routeHandler

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/gorilla/mux"
	"go.iondynamics.net/passPad/v1/passpad"
	"go.iondynamics.net/passPad/v1/passpad/lockdown"
	"go.iondynamics.net/passPad/v1/template"
	"go.iondynamics.net/webapp"
	"net/http"
	"net/mail"
)

func IndexGet(w http.ResponseWriter, req *http.Request) *webapp.Error {
	if acc := ensureAuth(w, req); acc != nil {
		if acc.ValidSecret {
			vaults, err := passpad.ListVaults(acc)
			if err != nil {
				return webapp.Write(err, err.Error(), http.StatusInternalServerError)
			}
			return template.Execute(w, acc, "indexGet", vaults)
		}
		http.Redirect(w, req, "/v1/setup", http.StatusFound)
	}
	return nil
}

func LoginGet(w http.ResponseWriter, req *http.Request) *webapp.Error {
	referrer := req.Header.Get("Referer")
	if referrer == "" {
		referrer = req.FormValue("referrer")
	}
	return template.Execute(w, nil, "loginGet", map[string]string{
		"HtmlTitle":    "Login PassPad",
		"FlashMessage": "",
		"Action":       "/v1/login",
		"RedirectTo":   referrer,
	})
}

func LoginPost(w http.ResponseWriter, req *http.Request) *webapp.Error {
	formUser := req.FormValue("input-user")
	formPass := req.FormValue("input-password")
	formToken := req.FormValue("input-token")
	redirectTo := req.FormValue("redirect-to")

	if formUser == "" || formPass == "" || formToken == "" {
		return template.Execute(w, nil, "loginGet", map[string]string{
			"HtmlTitle":    "Login PassPad",
			"FlashMessage": "Bitte geben Sie Nutzername, Passwort und das aktuelle Token an!",
			"Action":       "/v1/login",
		})
	} else {
		if !lockdown.IsLocked(formUser) {
			acc := passpad.AuthAccount(formUser, formPass)
			if acc != nil && passpad.ValidToken(acc, formToken) {
				session := sessions.GetSession(req)
				session.Set("user", formUser)
				session.Set("pass", formPass)
				if redirectTo == "" || redirectTo == "/v1/logout" {
					http.Redirect(w, req, "/", http.StatusFound)
				} else {
					http.Redirect(w, req, redirectTo, http.StatusFound)
				}
			} else {
				lockdown.Fail(formUser)
				return template.Execute(w, acc, "loginGet", map[string]string{
					"HtmlTitle":    "Login PassPad",
					"FlashMessage": "Nutzername, Passwort oder Token ungültig",
					"Action":       "/v1/login",
				})
			}
		} else {
			return template.Execute(w, nil, "loginGet", map[string]string{
				"HtmlTitle":    "Login PassPad",
				"FlashMessage": "Konto temporär gesperrt",
				"Action":       "/v1/login",
			})
		}
	}
	return nil
}

func LogoutGet(w http.ResponseWriter, req *http.Request) *webapp.Error {
	session := sessions.GetSession(req)
	session.Clear()
	http.Redirect(w, req, "/v1/login", http.StatusFound)
	return nil
}

func RegisterGet(w http.ResponseWriter, req *http.Request) *webapp.Error {
	if acc := ensureAuthNoRedirect(w, req); acc == nil {
		template.Execute(w, nil, "registerGet", nil)
	}
	return nil
}
func RegisterPost(w http.ResponseWriter, req *http.Request) *webapp.Error {
	formUser := req.FormValue("input-user")
	formPass := req.FormValue("input-password")
	formPass2 := req.FormValue("input-password2")
	redirectTo := req.FormValue("redirect-to")

	if formUser == "" || formPass == "" || formPass2 == "" {
		return template.Execute(w, nil, "loginGet", map[string]string{
			"HtmlTitle":    "Registrierung PassPad",
			"FlashMessage": "Bitte geben Sie Nutzername und Passwort an!",
			"Action":       "/v1/register",
		})
	} else {
		if formPass != formPass2 {
			return template.Execute(w, nil, "loginGet", map[string]string{
				"HtmlTitle":    "Registrierung PassPad",
				"FlashMessage": "Passwörter stimmen nicht überein",
				"Action":       "/v1/register",
			})
		}

		_, err := mail.ParseAddress(formUser)

		if err != nil {
			return template.Execute(w, nil, "loginGet", map[string]string{
				"HtmlTitle":    "Registrierung PassPad",
				"FlashMessage": "Nutzername ist keine gültige E-Mailadresse",
				"Action":       "/v1/register",
			})
		}

		if passpad.AccountExists(formUser) {
			return template.Execute(w, nil, "loginGet", map[string]string{
				"HtmlTitle":    "Registrierung PassPad",
				"FlashMessage": "Konto bereits vorhanden",
				"Action":       "/v1/register",
			})
		}

		if err = passpad.RegisterAccount(formUser, formPass); err != nil {
			return webapp.Write(err, "Couldn't register account", http.StatusInternalServerError)
		}

		session := sessions.GetSession(req)
		session.Set("user", formUser)
		session.Set("pass", formPass)
		if redirectTo == "" || redirectTo == "/v1/logout" {
			http.Redirect(w, req, "/", http.StatusFound)
		} else {
			http.Redirect(w, req, redirectTo, http.StatusFound)
		}
	}

	return nil
}

func SetupGet(w http.ResponseWriter, req *http.Request) *webapp.Error {
	if acc := ensureAuth(w, req); acc != nil {
		base64png, err := passpad.AccountSetup(acc)
		if err != nil {
			return webapp.Write(err, err.Error(), http.StatusInternalServerError)
		}
		return template.Execute(w, nil, "setupGet", map[string]string{
			"HtmlTitle": "Setup PassPad",
			"Action":    "/v1/setup",
			"Png":       base64png,
		})
	}
	return nil
}

func SetupPost(w http.ResponseWriter, req *http.Request) *webapp.Error {
	if acc := ensureAuth(w, req); acc != nil {
		formToken := req.FormValue("input-token")
		err := passpad.ValidateAccount(acc, formToken)
		if err != nil {
			http.Redirect(w, req, "/v1/setup", http.StatusFound)
		} else {
			http.Redirect(w, req, "/", http.StatusFound)
		}
	}

	return nil
}

func EntryGet(w http.ResponseWriter, req *http.Request) *webapp.Error {
	if acc := ensureAuth(w, req); acc != nil {
		identifier := mux.Vars(req)["identifier"]
		v, err := passpad.OpenVault(acc, identifier)
		if err != nil {
			return webapp.Write(err, err.Error(), http.StatusForbidden)
		}
		return template.Execute(w, acc, "entryGet", v)

	}
	return nil
}

func EntryPost(w http.ResponseWriter, req *http.Request) *webapp.Error {
	if acc := ensureAuth(w, req); acc != nil {
		identifier := mux.Vars(req)["identifier"]

		formName := req.FormValue("form-name")
		formUser := req.FormValue("form-user")
		formPass := req.FormValue("form-pass")
		formUrl := req.FormValue("form-url")

		if formName == "" {
			http.Redirect(w, req, "/v1/vault/"+identifier, http.StatusFound)
			return nil
		}

		err := passpad.UpsertEntry(acc, identifier, formName, formUser, formPass, formUrl, make(map[string]string))
		if err != nil {
			return webapp.Write(err, err.Error(), http.StatusForbidden)
		}
		http.Redirect(w, req, "/v1/vault/"+identifier, http.StatusFound)
	}

	return nil
}

func VaultPost(w http.ResponseWriter, req *http.Request) *webapp.Error {
	description := req.FormValue("form-description")
	if description == "" {
		http.Redirect(w, req, "/", http.StatusFound)
	}
	if acc := ensureAuth(w, req); acc != nil {
		identifier := req.FormValue("identifier")
		err := passpad.UpsertVault(acc, identifier, description)
		if err != nil {
			return webapp.Write(err, err.Error(), http.StatusForbidden)
		}
		http.Redirect(w, req, "/", http.StatusFound)
	}

	return nil
}
