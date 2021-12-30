package handler

import "net/http"

// FileUpload defined TODO
func (h *Handler) FileUpload(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(`{"id": 10}`))
}

// FileDownload defined TODO
func (h *Handler) FileDownload(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(`{"id": 10}`))
}
