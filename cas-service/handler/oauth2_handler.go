package handler

import (
	"comm/auth"
	"comm/logger"
	"net/http"
	"net/url"
	"os"

	"github.com/go-session/session"
)

func (h *Handler) OAuth2Authorize(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do OAuth2Authorize", acc.Name)
	}

	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = h.OAuth.HandleAuthorizeRequest(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

func (h *Handler) OAuth2Token(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do OAuth2Token", acc.Name)
	}
	err := h.OAuth.HandleTokenRequest(rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) OAuth2Login(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do OAuth2Login", acc.Name)
	}

	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		if r.Form == nil {
			if err := r.ParseForm(); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		store.Set("LoggedInUserID", r.Form.Get("username"))
		store.Save()
		rw.Header().Set("Location", "/cas/oauth2/affirm")
		rw.WriteHeader(http.StatusFound)
		return
	}
	outputHTML(rw, r, "static/login.html")
}

func (h *Handler) OAuth2Affirm(rw http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		rw.Header().Set("Location", "/login")
		rw.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(rw, r, "static/affirm.html")
}

func (h *Handler) UserAuthorizeHandler(rw http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(r.Context(), rw, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		rw.Header().Set("Location", "/cas/oauth2/login")
		rw.WriteHeader(http.StatusFound)
		return
	}
	// check
	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
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
