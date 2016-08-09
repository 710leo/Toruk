package render

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/unrolled/render"

	"github.com/710leo/Toruk/http/helper"
)

var Render *render.Render

var funcMap = template.FuncMap{
	"Example": helper.Example,
}

func Init() {
	debug := true
	Render = render.New(render.Options{
		Directory:     "views",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"{{", "}}"},
		Funcs:         []template.FuncMap{funcMap},
		IndentJSON:    false,
		IsDevelopment: debug,
	})
}

func Data(r *http.Request, key string, val interface{}) {
	m, ok := context.GetOk(r, "DATA_MAP")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		context.Set(r, "DATA_MAP", mm)
	} else {
		context.Set(r, "DATA_MAP", map[string]interface{}{key: val})
	}
}

func HTML(r *http.Request, w http.ResponseWriter, name string, htmlOpt ...render.HTMLOptions) {
	Render.HTML(w, http.StatusOK, name, context.Get(r, "DATA_MAP"), htmlOpt...)
}

func JSON(w http.ResponseWriter, v interface{}, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	Render.JSON(w, code, v)
}

func AutoJSON(w http.ResponseWriter, err error, v ...interface{}) {
	if err != nil {
		JSON(w, map[string]interface{}{"msg": err.Error()})
		return
	}

	if len(v) > 0 {
		JSON(w, map[string]interface{}{"msg": "", "data": v[0]})
	} else {
		JSON(w, map[string]interface{}{"msg": ""})
	}
}

func Text(w http.ResponseWriter, v string) {
	Render.Text(w, http.StatusOK, v)
}
