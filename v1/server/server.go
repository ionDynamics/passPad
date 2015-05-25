package server

import (
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"net"
	"net/http"
	"net/http/fcgi"

	idl "go.iondynamics.net/iDlogger"
	"go.iondynamics.net/iDnegroniLog"

	"go.iondynamics.net/passPad/v1/config"
	"go.iondynamics.net/passPad/v1/router"
)

func preflight(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)
}

func Listen() {
	logger := iDnegroniLog.NewMiddleware(idl.StandardLogger())

	logger.Stack2Http = config.Std.PassPad.DevelopmentEnv

	n := negroni.New(logger, negroni.NewStatic(http.Dir("public")))

	cookiestore := cookiestore.New([]byte(config.Std.Http.CookieSecret))
	cookiestore.Options(sessions.Options{MaxAge: config.Std.Http.CookieTimeout, Secure: config.Std.Http.CookieSecure})
	n.Use(sessions.Sessions("id_padpass_session", cookiestore))
	n.Use(negroni.HandlerFunc(preflight))

	n.UseHandler(router.New())

	if config.Std.Http.Fcgi {
		listener, err := net.Listen("tcp", config.Std.Http.Listen)
		if err != nil {
			idl.Fatal(err)
		}
		fcgi.Serve(listener, n)
	} else {
		n.Run(config.Std.Http.Listen)
	}
}
