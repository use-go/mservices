package main

import (
	"cas-service/handler"
	"comm/auth"
	"comm/define"
	"comm/logger"
	"comm/service/web"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	sAuth "github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/util/auth/token"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"golang.org/x/oauth2"
)

func main() {
	srv := web.New(web.Name("cas"))

	// create oauth2 server
	manager := manage.NewDefaultManager()
	// adapter grpc apigateway jwt check
	manager.MapAccessGenerate(auth.NewJWTAccessGenerate(
		token.WithPublicKey(sAuth.DefaultAuth.Options().PublicKey),
		token.WithPrivateKey(sAuth.DefaultAuth.Options().PrivateKey),
	))
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)
	auth := server.NewDefaultServer(manager)
	hdl := handler.Handler{
		OAuth: auth,
	}

	auth.SetAllowGetAccessRequest(true)
	auth.SetClientInfoHandler(server.ClientFormHandler)
	auth.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logger.Errorf(define.TODO, "Internal Error:%v", err.Error())
		return
	})
	auth.SetResponseErrorHandler(func(re *errors.Response) {
		logger.Errorf(define.TODO, "Response Error:%v", re.Error.Error())
	})
	auth.SetUserAuthorizationHandler(hdl.UserAuthorizeHandler)

	srv.HandleFunc("/oauth2/authorize", hdl.OAuth2Authorize)
	srv.HandleFunc("/oauth2/token", hdl.OAuth2Token)
	srv.HandleFunc("/oauth2/login", hdl.OAuth2Login)
	srv.HandleFunc("/oauth2/affirm", hdl.OAuth2Affirm)
	srv.HandleFunc("/account/add", hdl.AccountAdd)
	srv.HandleFunc("/account/del", hdl.AccountDel)
	srv.HandleFunc("/account/update", hdl.AccountUpdate)
	srv.HandleFunc("/account/list", hdl.AccountList)

	// For frontend test
	// Remove it in production
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://127.0.0.1:8888",
	})
	var (
		authServerURL = "http://127.0.0.1:8080"
		config        = oauth2.Config{
			ClientID:     "000000",
			ClientSecret: "999999",
			Scopes:       []string{"all"},
			RedirectURL:  "http://127.0.0.1:8080/cas/oauth2/adduser",
			Endpoint: oauth2.Endpoint{
				AuthURL:  authServerURL + "/cas/oauth2/authorize",
				TokenURL: authServerURL + "/cas/oauth2/token",
			},
		}
		genCodeChallengeS256 = func(s string) string {
			s256 := sha256.Sum256([]byte(s))
			return base64.URLEncoding.EncodeToString(s256[:])
		}
	)
	srv.HandleFunc("/oauth2/redirect", func(rw http.ResponseWriter, r *http.Request) {
		v := url.Values{
			"response_type":         {"code"},
			"state":                 {"xyz"},
			"redirect_uri":          {config.RedirectURL},
			"code_challenge":        {genCodeChallengeS256("s256example")},
			"code_challenge_method": {"S256"},
			"scope":                 {strings.Join(config.Scopes, " ")},
			"client_id":             {config.ClientID},
		}
		http.Redirect(rw, r, config.Endpoint.AuthURL+"?"+v.Encode(), http.StatusFound)
	})
	srv.HandleFunc("/oauth2/adduser", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		state := r.Form.Get("state")
		if state != "xyz" {
			http.Error(rw, "State invalid", http.StatusBadRequest)
			return
		}
		code := r.Form.Get("code")
		if code == "" {
			http.Error(rw, "Code not found", http.StatusBadRequest)
			return
		}
		v := url.Values{
			"grant_type":    {"authorization_code"},
			"code_verifier": {"s256example"},
			"client_id":     {config.ClientID},
			"client_secret": {config.ClientSecret},
			"redirect_uri":  {config.RedirectURL},
			"code":          {code},
		}
		rsp, err := http.Post(config.Endpoint.TokenURL, "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		body, err := ioutil.ReadAll(io.LimitReader(rsp.Body, 1<<20))
		rsp.Body.Close()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		if code := rsp.StatusCode; code < 200 || code > 299 {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(body)
	})

	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
