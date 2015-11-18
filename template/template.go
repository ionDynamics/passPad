package template

import (
	"html/template"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"go.iondynamics.net/iDhelper/randGen"
	"go.iondynamics.net/templice"
	"go.iondynamics.net/webapp"

	"go.iondynamics.net/passPad/account"
)

var tpl *templice.Template

func Load() error {
	tpl = templice.New(rice.MustFindBox("files"))

	funcMap := template.FuncMap{
		"rand": randGen.String,
	}

	tpl.SetPrep(func(templ *template.Template) *template.Template {
		return templ.Funcs(funcMap)
	})

	return tpl.Load()
}

func Execute(w http.ResponseWriter, acc *account.Account, tmpl string, data interface{}) *webapp.Error {
	err := tpl.ExecuteTemplate(w, tmpl+".tpl", data)
	if err != nil {
		return webapp.New(err, "Error executing template", http.StatusInternalServerError)
	}

	return nil
}
