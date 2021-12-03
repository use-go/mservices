package service

import (
	"comm/trace"
	"context"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/server"
)

var (
	debugWrapper = service.WrapHandler(func(call server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			defer trace.Debug(ctx, req.Endpoint(), req.Body(), rsp)()
			return call(ctx, req, rsp)
		}
	})
)
