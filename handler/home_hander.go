package handler

import (
	"net/http"

	"github.com/qinyening/Toruk/http/render"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	render.HTML(r, w, "home/index")
}
