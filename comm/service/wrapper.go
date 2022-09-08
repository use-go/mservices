package service

import (
	"comm/trace"
	"context"

	"github.com/2637309949/micro/v3/service"
	"github.com/2637309949/micro/v3/service/client"
	"github.com/2637309949/micro/v3/service/context/metadata"
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

type clientWrapper struct {
	client.Client
	headers metadata.Metadata
}

var (
	HeaderPrefix = "Micro-"
)

func (c *clientWrapper) setHeaders(ctx context.Context) context.Context {
	// copy metadata
	mda, _ := metadata.FromContext(ctx)
	md := metadata.Copy(mda)

	// set headers
	for k, v := range c.headers {
		if _, ok := md[k]; !ok {
			md[k] = v
		}
	}

	return metadata.NewContext(ctx, md)
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	ctx = c.setHeaders(ctx)
	return c.Client.Call(ctx, req, rsp, opts...)
}

func (c *clientWrapper) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	ctx = c.setHeaders(ctx)
	return c.Client.Stream(ctx, req, opts...)
}

func (c *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	ctx = c.setHeaders(ctx)
	return c.Client.Publish(ctx, p, opts...)
}

// FromService wraps a client to inject From-Service header into metadata
func FromService(name string) func(client.Client) client.Client {
	return func(c client.Client) client.Client {
		return &clientWrapper{
			c,
			metadata.Metadata{
				HeaderPrefix + "From-Service": name,
			},
		}
	}
}
