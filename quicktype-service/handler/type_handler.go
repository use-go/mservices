package handler

import (
	whttp "comm/service/web/http"
	"net/http"
)

// FileUpload defined TODO
func (h *Handler) Type(rw http.ResponseWriter, r *http.Request) {
	whttp.OutputHTML(rw, r, "static/index.html")
}
