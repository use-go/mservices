package http

import (
	"github.com/2637309949/micro/v3/service/api/router"
	"github.com/2637309949/micro/v3/service/registry"
)

type Options struct {
	s        *registry.Service
	Router   router.Router
	Registry registry.Registry
}

type Option func(*Options)

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func WithService(s *registry.Service) Option {
	return func(o *Options) {
		o.s = s
	}
}
