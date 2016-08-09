package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/710leo/Toruk/handler"
)

func ConfigRouter(r *mux.Router) {
	configConfRoutes(r)
}

func configConfRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	}).Methods("GET")

	r.HandleFunc("/home", handler.HomeIndex).Methods("GET")
}
