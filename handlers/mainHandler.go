// Package handlers provides needed handlers for the web app.
package handlers

import (
	"log"
	"net/http"
	"text/template"

	s "github.com/mellenooijen/gowebapp/structs"
)

var ExampleUser = &s.User{Id: 1, Name: "Melle"}

// MainHandler is the basic handler used for all pages that don't require any intervention.
// It renders the given page (or index when r.URL.Path is /), and exposes a User object and page title string to the template found in ./html/(title).html.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := r.URL.Path[1:]
	if tmpl == "" {
		tmpl = "index"
	}
	t, e := template.ParseFiles("html/" + tmpl + ".html")
	if e != nil {
		log.Print(e)
		ErrorHandler(w, r, s.IntError{Code: 404, Page: r.URL.Path, Output: e.Error()})
	} else {
		d := s.PageData{Title: tmpl, User: ExampleUser.Name}
		t.Execute(w, d)
	}
}
