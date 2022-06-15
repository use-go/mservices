package handler

import (
	"context"
	"time"

	"comm/auth"
	"comm/errors"
	"comm/logger"
	"comm/mark"

	"proto/cache"
)

func (h *Handler) Get(ctx context.Context, req *cache.GetRequest, rsp *cache.GetResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Get")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Get", acc.Name)
	}
	err = h.CacheStore.Get(req.Key, &rsp.Value)
	timemark.Mark("Get")
	if err != nil {
		return errors.InternalServerError("CacheStore.Get failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Set(ctx context.Context, req *cache.SetRequest, rsp *cache.SetResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Set")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Set", acc.Name)
	}
	err = h.CacheStore.Set(req.Key, req.Value, time.Duration(req.Expire))
	timemark.Mark("Set")
	if err != nil {
		return errors.InternalServerError("CacheStore.Set failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Add(ctx context.Context, req *cache.AddRequest, rsp *cache.AddResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Add")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Add", acc.Name)
	}
	err = h.CacheStore.Add(req.Key, req.Value, time.Duration(req.Expire))
	timemark.Mark("Add")
	if err != nil {
		return errors.InternalServerError("CacheStore.Add failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Replace(ctx context.Context, req *cache.ReplaceRequest, rsp *cache.ReplaceResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Replace")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Replace", acc.Name)
	}
	err = h.CacheStore.Replace(req.Key, req.Value, time.Duration(req.Expire))
	timemark.Mark("Replace")
	if err != nil {
		return errors.InternalServerError("CacheStore.Replace failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *cache.DeleteRequest, rsp *cache.DeleteResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Delete")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Delete", acc.Name)
	}
	err = h.CacheStore.Delete(req.Key)
	timemark.Mark("Delete")
	if err != nil {
		return errors.InternalServerError("CacheStore.Delete failed %v", err.Error())
	}
	return nil
}

func (h *Handler) Increment(ctx context.Context, req *cache.IncrementRequest, rsp *cache.IncrementResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Increment")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Increment", acc.Name)
	}
	inc, err := h.CacheStore.Increment(req.Key, uint64(req.Value))
	timemark.Mark("Increment")
	if err != nil {
		return errors.InternalServerError("CacheStore.Increment failed %v", err.Error())
	}
	rsp.Value = int64(inc)
	return nil
}

func (h *Handler) Decrement(ctx context.Context, req *cache.DecrementRequest, rsp *cache.DecrementResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Decrement")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Decrement", acc.Name)
	}
	desc, err := h.CacheStore.Decrement(req.Key, uint64(req.Value))
	timemark.Mark("Decrement")
	if err != nil {
		return errors.InternalServerError("CacheStore.Decrement failed %v", err.Error())
	}
	rsp.Value = int64(desc)
	return nil
}

func (h *Handler) Flush(ctx context.Context, req *cache.FlushRequest, rsp *cache.FlushResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "Flush")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Flush", acc.Name)
	}
	err = h.CacheStore.Flush()
	timemark.Mark("Flush")
	if err != nil {
		return errors.InternalServerError("CacheStore.Flush failed %v", err.Error())
	}
	return nil
}
