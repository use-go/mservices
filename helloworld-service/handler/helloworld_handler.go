package handler

import (
	"comm/auth"
	"comm/errors"
	"comm/logger"
	"context"
	"fmt"
	"proto/helloworld"
	"time"
)

// DeleteInfo defined TODO
func (h *Handler) DeleteInfo(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteInfo", acc.Name)
	}
	err := h.CacheStore.Set(ctx, "DeleteInfo", 1, 10*time.Second)
	if err != nil {
		return errors.InternalServerError("set failed %v", err)
	}
	_, err = h.CacheStore.Increment(ctx, "DeleteInfo", 2)
	if err != nil {
		return errors.InternalServerError("increment failed %v", err)
	}
	ret, err := h.CacheStore.Decrement(ctx, "DeleteInfo", 1)
	if err != nil {
		return errors.InternalServerError("decrement failed %v", err)
	}
	rsp.Name = "Hello " + req.Name + fmt.Sprintf("%v", ret)
	return nil
}

// UpdateInfo defined TODO
func (h *Handler) UpdateInfo(ctx context.Context, req *helloworld.Info, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do UpdateInfo", acc.Name)
	}
	rsp.Name = "Hello " + req.Name
	return nil
}

// InsertInfo defined TODO
func (h *Handler) InsertInfo(ctx context.Context, req *helloworld.Info, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do InsertInfo", acc.Name)
	}
	rsp.Name = "Hello " + req.Name
	return nil
}

// QueryInfoDetail defined TODO
func (h *Handler) QueryInfoDetail(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfoDetail", acc.Name)
	}
	rsp.Name = "Hello " + req.Name
	return nil
}

// QueryInfo defined TODO
func (h *Handler) QueryInfo(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.InfoList) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfo", acc.Name)
	}
	rsp.Data = append(rsp.Data, &helloworld.Info{})
	return nil
}
