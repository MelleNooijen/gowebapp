// Package server implements the basics for running the server and routing to handlers.
package server

import (
	"log"
	"net/http"

	"github.com/mellenooijen/gowebapp/handlers"
)

func Run() {
	http.HandleFunc("/direct/", handlers.FileHandler)
	http.HandleFunc("/", handlers.MainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
