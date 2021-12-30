package web

import (
	"context"
	"time"

	"github.com/google/uuid"
)

var (
	// For serving
	Namespace      = "micro"
	DefaultName    = "go-web"
	DefaultVersion = "latest"
	DefaultId      = uuid.New().String()
	DefaultAddress = ":0"

	// for registration
	DefaultRegisterTTL      = time.Minute
	DefaultRegisterInterval = time.Second * 30
	DefaultRegisterCheck    = func(context.Context) error { return nil }
	// static directory
	DefaultStaticDir = "html"
)

type Option func(o *Options)

// NewService returns a new web.Service
func New(opts ...Option) *Service {
	opts = append(opts, Wrapper(debugWrapper))
	return newService(opts...)
}
