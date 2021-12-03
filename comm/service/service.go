package service

import (
	"time"

	"github.com/micro/micro/v3/service"
)

func New(opts ...service.Option) *service.Service {
	opts = append(opts, service.Version("latest"))
	opts = append(opts, service.RegisterTTL(5*time.Second))
	opts = append(opts, service.RegisterInterval(5*time.Second))
	opts = append(opts, debugWrapper)
	return service.New(opts...)
}

func Name(n string) service.Option {
	return service.Name(n)
}
