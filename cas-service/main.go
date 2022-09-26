package main

import (
	"cas-service/handler"
	"cas-service/util/cookie"
	"comm/auth"
	"comm/define"
	"comm/logger"
	"comm/service/web"
	"encoding/gob"
	"net/url"

	"github.com/2637309949/micro/v3/util/auth/token"
	"github.com/2637309949/micro/v3/util/user"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session/v3"
)

const (
	hashKey = "x8hsd7sdbvs938b2d790h"
)

func main() {
	// Create service
	srv := web.New(web.Name("cas"))

	// Create oauth2 server
	manager := manage.NewDefaultManager()
	// Adapter grpc apigateway jwt check
	privKey, pubKey, _ := user.GetJWTCerts()
	manager.MapAccessGenerate(auth.NewJWTAccessGenerate(token.WithPublicKey(string(pubKey)), token.WithPrivateKey(string(privKey))))
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	auth := server.NewDefaultServer(manager)
	hdl := handler.Handler{
		OAuthServer: auth,
		ClientStore: clientStore,
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

	// Third Party Authorization
	srv.HandleFunc("/oauth2/authorize", hdl.OAuth2Authorize)
	srv.HandleFunc("/oauth2/token", hdl.OAuth2Token)
	srv.HandleFunc("/oauth2/affirm", hdl.OAuth2Affirm)

	// The user authorization
	srv.HandleFunc("/user/login", hdl.UserLogin)
	srv.HandleFunc("/user/logout", hdl.UserLogout)

	// Implicit Grant​, Client Authorization
	srv.HandleFunc("/client/authorize", hdl.ClientAuthorize)
	srv.HandleFunc("/client/token", hdl.ClientToken)

	// Session store
	gob.Register(url.Values{})
	session.InitManager(
		session.SetCookieName("session_id"),
		session.SetStore(
			cookie.NewCookieStore(
				cookie.SetCookieName("store"),
				cookie.SetHashKey([]byte(hashKey)),
			),
		),
	)

	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
