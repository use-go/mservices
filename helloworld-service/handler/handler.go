package handler

import (
	"comm/store"
	"proto/email"
)

type Handler struct {
	CacheService store.Cache
	EmailService email.EmailService
}
