package handler

import (
	"context"
	"time"

	"comm/auth"
	"comm/errors"
	"comm/logger"

	"proto/cache"
)

func (h *Handler) Get(ctx context.Context, req *cache.GetReq, rsp *cache.GetRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Get", acc.Name)
	}
	err := h.CacheStore.Get(req.Key, &rsp.Value)
	if err != nil {
		return errors.InternalServerError("CacheStore.Get failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Set(ctx context.Context, req *cache.SetReq, rsp *cache.SetRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Set", acc.Name)
	}
	err := h.CacheStore.Set(req.Key, req.Value, time.Duration(req.Expire))
	if err != nil {
		return errors.InternalServerError("CacheStore.Set failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Add(ctx context.Context, req *cache.AddReq, rsp *cache.AddRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Add", acc.Name)
	}
	err := h.CacheStore.Add(req.Key, req.Value, time.Duration(req.Expire))
	if err != nil {
		return errors.InternalServerError("CacheStore.Add failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Replace(ctx context.Context, req *cache.ReplaceReq, rsp *cache.ReplaceRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Replace", acc.Name)
	}
	err := h.CacheStore.Replace(req.Key, req.Value, time.Duration(req.Expire))
	if err != nil {
		return errors.InternalServerError("CacheStore.Replace failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *cache.DeleteReq, rsp *cache.DeleteRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Delete", acc.Name)
	}
	err := h.CacheStore.Delete(req.Key)
	if err != nil {
		return errors.InternalServerError("CacheStore.Delete failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Increment(ctx context.Context, req *cache.IncrementReq, rsp *cache.IncrementRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Increment", acc.Name)
	}
	inc, err := h.CacheStore.Increment(req.Key, uint64(req.Value))
	if err != nil {
		return errors.InternalServerError("CacheStore.Increment failed %v", err.Error())
	}
	rsp.Value = int64(inc)
	return nil
}

func (h *Handler) Decrement(ctx context.Context, req *cache.DecrementReq, rsp *cache.DecrementRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Decrement", acc.Name)
	}
	desc, err := h.CacheStore.Decrement(req.Key, uint64(req.Value))
	if err != nil {
		return errors.InternalServerError("CacheStore.Decrement failed %v", err.Error())
	}
	rsp.Value = int64(desc)
	return nil
}

func (h *Handler) Flush(ctx context.Context, req *cache.FlushReq, rsp *cache.FlushRes) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Flush", acc.Name)
	}
	err := h.CacheStore.Flush()
	if err != nil {
		return errors.InternalServerError("CacheStore.Flush failed %v", err.Error())
	}
	return nil
}
