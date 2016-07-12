package http

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/qinyening/Toruk/g"
	"github.com/qinyening/Toruk/http/cookie"
	"github.com/qinyening/Toruk/http/middleware"
	"github.com/qinyening/Toruk/http/render"
)

func Start() {
	render.Init()
	cookie.Init()

	r := mux.NewRouter().StrictSlash(false)
	ConfigRouter(r)

	n := negroni.New()
	n.Use(middleware.NewRecovery())
	n.UseHandler(r)
	n.Run(g.Config().Http.Listen)
}
