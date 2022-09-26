package handler

import (
	"comm/auth"
	"comm/logger"
	"comm/mark"
	whttp "comm/service/web/http"
	"net/http"

	"github.com/go-session/session/v3"
)

// UserLogin defined todo
func (h *Handler) UserLogin(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "UserLogin")()

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

	timemark.Mark("ParseForm")
	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// valid user idx
	timemark.Mark("ValidUser")
	email, username, password := r.Form.Get("email"), r.Form.Get("username"), r.Form.Get("password")
	_, _ = email, password
	store.Set("LoggedInUserID", username)
	store.Save()
	rw.Header().Set("Location", "/cas/oauth2/affirm")
	rw.WriteHeader(http.StatusFound)
}

// UserLogout defined todo
func (h *Handler) UserLogout(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "UserLogout")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do UserLogout", acc.Name)
	}

	timemark.Mark("Destroy")
	err := session.Destroy(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
