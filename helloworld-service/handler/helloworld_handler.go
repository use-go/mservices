package handler

import (
	"comm/errors"
	"comm/logger"
	"context"
	"proto/helloworld"
)

// DeleteHelloworld defined TODO
func (h *Handler) DeleteHelloworld(ctx context.Context, req *helloworld.HelloworldFilter, rsp *helloworld.Helloworld) error {
	logger.Infof("Do something")
	rsp.Name = "Hello " + req.Name
	return errors.New("helloworld.DeleteHelloworld", "test error", int32(50001))
}

// UpdateHelloworld defined TODO
func (h *Handler) UpdateHelloworld(ctx context.Context, req *helloworld.Helloworld, rsp *helloworld.Helloworld) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

// InsertHelloworld defined TODO
func (h *Handler) InsertHelloworld(ctx context.Context, req *helloworld.Helloworld, rsp *helloworld.Helloworld) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

// QueryHelloworldDetail defined TODO
func (h *Handler) QueryHelloworldDetail(ctx context.Context, req *helloworld.HelloworldFilter, rsp *helloworld.Helloworld) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

// QueryHelloworld defined TODO
func (h *Handler) QueryHelloworld(ctx context.Context, req *helloworld.HelloworldFilter, rsp *helloworld.HelloworldList) error {
	rsp.Data = append(rsp.Data, &helloworld.Helloworld{})
	return nil
}
