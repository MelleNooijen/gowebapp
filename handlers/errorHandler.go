package handlers

import (
	"net/http"
	"text/template"

	s "github.com/mellenooijen/gowebapp/structs"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, e s.IntError) {
	t, _ := template.ParseFiles("html/error.html") // use _ since no errors should be produced
	w.WriteHeader(e.Code)
	t.Execute(w, e)
}
