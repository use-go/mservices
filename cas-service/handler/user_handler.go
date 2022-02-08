package handler

import (
	"comm/auth"
	"comm/logger"
	whttp "comm/service/web/http"
	"net/http"

	"github.com/go-session/session/v3"
)

func (h *Handler) UserLogin(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do UserLogin", acc.Name)
	}

	if r.Method == "GET" {
		whttp.OutputHTML(rw, r, "static/login.html")
		return
	}

	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	userID := r.Form.Get("username")
	store.Set("LoggedInUserID", userID)
	store.Save()
	rw.Header().Set("Location", "/cas/oauth2/affirm")
	rw.WriteHeader(http.StatusFound)
}

func (h *Handler) UserLogout(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do UserLogout", acc.Name)
	}
	err := session.Destroy(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
