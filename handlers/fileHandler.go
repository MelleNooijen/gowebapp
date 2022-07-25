package handlers

import (
	"errors"
	"net/http"
	"os"

	s "github.com/mellenooijen/gowebapp/structs"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "./public/direct/" + string(r.URL.Path[len("/direct/"):])
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		ErrorHandler(w, r, s.IntError{Code: 404, Output: err.Error(), Page: filePath})
	} else {
		http.ServeFile(w, r, filePath)
	}
}
