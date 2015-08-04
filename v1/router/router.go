package router

import (
	"net/http"

	"github.com/gorilla/mux"
	handler "go.iondynamics.net/passPad/v1/routeHandler"
	"go.iondynamics.net/webapp"
)

func New() *mux.Router {
	return provision(mux.NewRouter().StrictSlash(true))
}

func provision(r *mux.Router) *mux.Router {
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/v1/", http.StatusFound)
	})

	v1GetHead := r.PathPrefix("/v1").Methods("GET", "HEAD").Subrouter()
	v1GetHead.Handle("/", webapp.Handler(handler.IndexGet))
	v1GetHead.Handle("/login", webapp.Handler(handler.LoginGet))
	v1GetHead.Handle("/register", webapp.Handler(handler.RegisterGet))
	v1GetHead.Handle("/setup", webapp.Handler(handler.SetupGet))
	v1GetHead.Handle("/logout", webapp.Handler(handler.LogoutGet))
	v1GetHead.Handle("/vault/{identifier}", webapp.Handler(handler.EntryGet))

	v1Post := r.PathPrefix("/v1").Methods("POST").Subrouter()
	v1Post.Handle("/login", webapp.Handler(handler.LoginPost))
	v1Post.Handle("/register", webapp.Handler(handler.RegisterPost))
	v1Post.Handle("/setup", webapp.Handler(handler.SetupPost))
	v1Post.Handle("/vault", webapp.Handler(handler.VaultPost))
	v1Post.Handle("/vault/{identifier}", webapp.Handler(handler.EntryPost))

	return r
}
