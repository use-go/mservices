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
	base, err := url.Parse(r.URL.Query().Get("redirect_uri"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	h.ClientStore.Set(base.Host, &models.Client{
		ID:     base.Host,
		Secret: fmt.Sprintf("#$!%v", base.Host),
		Domain: fmt.Sprintf("%v://%v", base.Scheme, base.Host),
	})
	v := url.Values{
		"response_type":         {"code"},
		"state":                 {r.URL.Query().Get("state")},
		"redirect_uri":          {r.URL.Query().Get("redirect_uri")},
		"code_challenge":        {codeChallengeS256("s256example")},
		"code_challenge_method": {"S256"},
		"scope":                 {strings.Join([]string{"all"}, " ")},
		"client_id":             {base.Host},
	}
	r.URL.RawQuery = v.Encode()
	h.OAuth2Authorize(rw, r)
}

func (h *Handler) ClientToken(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	base, err := url.Parse(r.Form.Get("redirect_uri"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	code := r.Form.Get("code")
	if code == "" {
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
		"client_id":     {base.Host},
		"client_secret": {cli.GetSecret()},
		"redirect_uri":  {r.Form.Get("redirect_uri")},
		"code":          {code},
	}
	r.Form = v
	h.OAuth2Token(rw, r)
}
