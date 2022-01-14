package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-oauth2/oauth2/v4/models"
)

func codeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}

func (h *Handler) ClientAuthorize(rw http.ResponseWriter, r *http.Request) {
	redirectUri, _ := url.QueryUnescape(r.URL.Query().Get("redirect_uri"))
	state, _ := url.QueryUnescape(r.URL.Query().Get("state"))
	base, err := url.Parse(redirectUri)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	h.ClientStore.Set(base.Host, &models.Client{
		ID:     base.Host,
		Secret: fmt.Sprintf("%v#4!", base.Host),
		Domain: fmt.Sprintf("%v://%v", base.Scheme, base.Host),
	})
	scope := []string{"all"}
	v := url.Values{
		"response_type":         {"code"},
		"state":                 {state},
		"redirect_uri":          {redirectUri},
		"code_challenge":        {codeChallengeS256("s256example")},
		"code_challenge_method": {"S256"},
		"scope":                 {strings.Join(scope, " ")},
		"client_id":             {base.Host},
	}
	r.URL.RawQuery = v.Encode()
	h.OAuth2Authorize(rw, r)
}

func (h *Handler) ClientToken(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	redirectUri, _ := url.QueryUnescape(r.Form.Get("redirect_uri"))
	if len(redirectUri) == 0 {
		http.Error(rw, "redirectUri not found", http.StatusBadRequest)
		return
	}
	base, err := url.Parse(redirectUri)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	code := r.Form.Get("code")
	if len(code) == 0 {
		http.Error(rw, "code not found", http.StatusBadRequest)
		return
	}

	cli, err := h.ClientStore.GetByID(r.Context(), base.Host)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	v := url.Values{
		"grant_type":    {"authorization_code"},
		"code_verifier": {"s256example"},
		"client_id":     {cli.GetID()},
		"client_secret": {cli.GetSecret()},
		"redirect_uri":  {redirectUri},
		"code":          {code},
	}
	r.Form = v
	h.OAuth2Token(rw, r)
}
