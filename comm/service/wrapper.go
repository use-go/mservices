package service

import (
	"comm/trace"
	"context"
	"fmt"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/server"
)

var (
	debugWrapper = service.WrapHandler(func(call server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			endpoint, body := req.Endpoint(), req.Body()
			fmt.Println("------", req.Header())
			defer trace.Debug(ctx, endpoint, body, rsp)()
			return call(ctx, req, rsp)
		}
	})
)
