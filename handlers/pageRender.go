package handlers

import (
	"net/http"
	"text/template"

	s "github.com/mellenooijen/gowebapp/structs"
)

func RenderPage(w http.ResponseWriter, r *http.Request, page string, d s.PageData) {
	t, e := template.ParseFiles("html/"+page+".html", "html/base.html")
	if e != nil {
		c := 404
		if e.Error()[:8] == "template" {
			c = 500
		}
		ErrorHandler(w, r, s.IntError{Code: c, Page: r.URL.Path, Output: e.Error()}, e)
	} else {
		e := t.ExecuteTemplate(w, "base", d)
		if e != nil {
			ErrorHandler(w, r, s.IntError{Code: 500, Page: r.URL.Path, Output: e.Error()}, e)
		}
	}
}
