package handlers

import (
	"log"
	"net/http"

	s "github.com/mellenooijen/gowebapp/structs"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, i s.IntError, e error) {
	log.Print(e)
	w.WriteHeader(i.Code)
	RenderPage(w, r, "error", s.PageData{Title: "Error", User: ExampleUser.Name, Error: i})
}
