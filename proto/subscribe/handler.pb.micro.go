// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/subscribe/handler.proto

package subscribe

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/2637309949/micro/v3/service/api"
	client "github.com/2637309949/micro/v3/service/client"
	server "github.com/2637309949/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Subscribe service

func NewSubscribeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Subscribe service

type SubscribeService interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (Subscribe_SubscribeService, error)
}

type subscribeService struct {
	c    client.Client
	name string
}

func NewSubscribeService(name string, c client.Client) SubscribeService {
	return &subscribeService{
		c:    c,
		name: name,
	}
}

func (c *subscribeService) Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error) {
	req := c.c.NewRequest(c.name, "Subscribe.Publish", in)
	out := new(PublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeService) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...client.CallOption) (Subscribe_SubscribeService, error) {
	req := c.c.NewRequest(c.name, "Subscribe.Subscribe", &SubscribeRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &subscribeServiceSubscribe{stream}, nil
}

type Subscribe_SubscribeService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*SubscribeResponse, error)
}

type subscribeServiceSubscribe struct {
	stream client.Stream
}

func (x *subscribeServiceSubscribe) Close() error {
	return x.stream.Close()
}

func (x *subscribeServiceSubscribe) Context() context.Context {
	return x.stream.Context()
}

func (x *subscribeServiceSubscribe) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *subscribeServiceSubscribe) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *subscribeServiceSubscribe) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Subscribe service

type SubscribeHandler interface {
	Publish(context.Context, *PublishRequest, *PublishResponse) error
	Subscribe(context.Context, *SubscribeRequest, Subscribe_SubscribeStream) error
}

func RegisterSubscribeHandler(s server.Server, hdlr SubscribeHandler, opts ...server.HandlerOption) error {
	type subscribe interface {
		Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error
		Subscribe(ctx context.Context, stream server.Stream) error
	}
	type Subscribe struct {
		subscribe
	}
	h := &subscribeHandler{hdlr}
	return s.Handle(s.NewHandler(&Subscribe{h}, opts...))
}

type subscribeHandler struct {
	SubscribeHandler
}

func (h *subscribeHandler) Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error {
	return h.SubscribeHandler.Publish(ctx, in, out)
}

func (h *subscribeHandler) Subscribe(ctx context.Context, stream server.Stream) error {
	m := new(SubscribeRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SubscribeHandler.Subscribe(ctx, m, &subscribeSubscribeStream{stream})
}

type Subscribe_SubscribeStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*SubscribeResponse) error
}

type subscribeSubscribeStream struct {
	stream server.Stream
}

func (x *subscribeSubscribeStream) Close() error {
	return x.stream.Close()
}

func (x *subscribeSubscribeStream) Context() context.Context {
	return x.stream.Context()
}

func (x *subscribeSubscribeStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *subscribeSubscribeStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *subscribeSubscribeStream) Send(m *SubscribeResponse) error {
	return x.stream.Send(m)
}
