package http

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/710leo/Toruk/g"
	"github.com/710leo/Toruk/http/cookie"
	"github.com/710leo/Toruk/http/middleware"
	"github.com/710leo/Toruk/http/render"
)

func Start() {
	render.Init()
	cookie.Init()

	r := mux.NewRouter().StrictSlash(false)
	ConfigRouter(r)

	n := negroni.New()

	if g.Config().Debug {
		n.Use(negroni.NewLogger())
	}
	n.Use(middleware.NewRecovery())
	n.UseHandler(r)
	n.Run(g.Config().Http.Listen)
}
