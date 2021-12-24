package http

import (
	"net/http"

	"github.com/2637309949/micro/v3/service/registry"
	"github.com/2637309949/micro/v3/util/selector/roundrobin"
)

func NewRoundTripper(opts ...Option) http.RoundTripper {
	options := Options{
		Registry: registry.DefaultRegistry,
	}
	for _, o := range opts {
		o(&options)
	}

	return &roundTripper{
		rt:   http.DefaultTransport,
		st:   roundrobin.NewSelector(),
		opts: options,
	}
}
