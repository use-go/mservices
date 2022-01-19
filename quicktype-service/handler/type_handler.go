package handler

import (
	"net/http"
	"os"
)

// FileUpload defined TODO
func (h *Handler) Type(rw http.ResponseWriter, r *http.Request) {
	outputHTML(rw, r, "static/index.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
