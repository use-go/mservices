package client

import (
	"github.com/2637309949/micro/v3/service/client"
)

func WithAuthToken() client.CallOption {
	return func(o *client.CallOptions) {
		o.AuthToken = true
	}
}
