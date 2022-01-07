package handler

import (
	"comm/auth"
	"comm/logger"
	"net/http"
)

// AccountAdd defined TODO
func (h *Handler) AccountAdd(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do AccountAdd", acc.Name)
	}
}

// AccountDel defined TODO
func (h *Handler) AccountDel(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do AccountDel", acc.Name)
	}
}

// AccountUpdate defined TODO
func (h *Handler) AccountUpdate(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do AccountUpdate", acc.Name)
	}
}

// AccountList defined TODO
func (h *Handler) AccountList(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do AccountList", acc.Name)
	}
}
