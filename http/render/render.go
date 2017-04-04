package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/unrolled/render"

	"github.com/710leo/Toruk/g"
	"github.com/710leo/Toruk/http/helper"
)

var Render *render.Render

var funcMap = template.FuncMap{
	"Example": helper.Example,
}

func Init() {
	Render = render.New(render.Options{
		Directory:     "views",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"{{", "}}"},
		Funcs:         []template.FuncMap{funcMap},
		IndentJSON:    false,
		IsDevelopment: g.Config().Debug,
	})
}

func Put(r *http.Request, key string, val interface{}) {
	m, ok := context.GetOk(r, "_DATA_MAP_")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		context.Set(r, "_DATA_MAP_", mm)
	} else {
		context.Set(r, "_DATA_MAP_", map[string]interface{}{key: val})
	}
}

func HTML(r *http.Request, w http.ResponseWriter, name string, htmlOpt ...render.HTMLOptions) {
	Render.HTML(w, http.StatusOK, name, context.Get(r, "_DATA_MAP_"), htmlOpt...)
}

func Text(w http.ResponseWriter, v string, codes ...int) {
	code := http.StatusOK
	if len(codes) > 0 {
		code = codes[0]
	}
	Render.Text(w, code, v)
}

func Error(w http.ResponseWriter, err error) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}

	Render.JSON(w, http.StatusOK, map[string]string{"msg": msg})
}

func Message(w http.ResponseWriter, format string, args ...interface{}) {
	Render.JSON(w, http.StatusOK, map[string]string{"msg": fmt.Sprintf(format, args...)})
}

func Data(w http.ResponseWriter, v interface{}, msg ...string) {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}

	Render.JSON(w, http.StatusOK, map[string]interface{}{"msg": m, "data": v})
}
