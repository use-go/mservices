package handler

import (
	"comm/auth"
	"comm/errors"
	"comm/logger"
	"context"
	"proto/helloworld"
)

// DeleteInfo defined TODO
func (h *Handler) DeleteInfo(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteInfo", acc.Name)
	}
	rsp.Name = "Hello " + req.Name
	return errors.New("test error", int32(50001))
}

// UpdateInfo defined TODO
func (h *Handler) UpdateInfo(ctx context.Context, req *helloworld.Info, rsp *helloworld.Info) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

// InsertInfo defined TODO
func (h *Handler) InsertInfo(ctx context.Context, req *helloworld.Info, rsp *helloworld.Info) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

// QueryInfoDetail defined TODO
func (h *Handler) QueryInfoDetail(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.Info) error {
	rsp.Name = "Hello " + req.Name
	return nil
}

// QueryInfo defined TODO
func (h *Handler) QueryInfo(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.InfoList) error {
	rsp.Data = append(rsp.Data, &helloworld.Info{})
	return nil
}
