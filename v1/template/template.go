package template

import (
	"html/template"
	"net/http"
	"sync"
	"time"

	"go.iondynamics.net/iDhelper/randGen"
	"go.iondynamics.net/passPad/v1/passpad/account"
	"go.iondynamics.net/webapp"
)

var templates *template.Template
var tplReloadMutex sync.RWMutex

func Execute(w http.ResponseWriter, acc *account.Account, tmpl string, data interface{}) *webapp.Error {
	tplReloadMutex.RLock()
	defer tplReloadMutex.RUnlock()
	err := templates.ExecuteTemplate(w, tmpl+".tpl", data)

	if err != nil {
		return webapp.New(err, "Error executing template", http.StatusInternalServerError)
	} else {
		return nil
	}
}

func Load(glob string) {
	tplReloadMutex.Lock()
	defer tplReloadMutex.Unlock()
	funcMap := template.FuncMap{
		"rand": randGen.String,
	}
	templates = template.Must(template.New("main").Funcs(funcMap).ParseGlob(glob))
	time.Sleep(1 * time.Second)
}

func getNavCat(tmpl string) string {
	ret := ""
	switch tmpl {
	case "indexGet":
		ret = "home"
	case "contact":
		ret = "contact"
	default:
		ret = ""
	}

	return ret
}
