package handler

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

type Handler struct {
	OAuthServer *server.Server
	ClientStore *store.ClientStore
}
