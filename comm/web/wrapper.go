package web

import (
	"comm/trace"
	"context"

	"github.com/2637309949/micro/v3/service"
	"github.com/2637309949/micro/v3/service/server"
)

var (
	debugWrapper = service.WrapHandler(func(call server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			endpoint, body := req.Endpoint(), req.Body()
			defer trace.Debug(ctx, endpoint, body, rsp)()
			return call(ctx, req, rsp)
		}
	})
)
