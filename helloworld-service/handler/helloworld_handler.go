package handler

import (
	"comm/auth"
	"comm/db"
	"comm/errors"
	"comm/logger"
	"context"
	"helloworld-service/model"
	"proto/helloworld"

	"github.com/jinzhu/copier"
)

// DeleteInfo defined TODO
func (h *Handler) DeleteInfo(ctx context.Context, req *helloworld.InfoFilter, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteInfo", acc.Name)
	}

	if len(req.Name) == 0 {
		return errors.BadRequest("Name required")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	where := model.Info{
		Name: req.GetName(),
	}
	err = h.DeleteInfoDB(ctx, session, &where)
	if err != nil {
		logger.Errorf(ctx, "DeleteInfoDB failed. [%v]", err)
		return errors.InternalServerError("deleteInfoDB failed %v", err.Error())
	}
	rsp.Name = where.Name
	return nil
}

// UpdateInfo defined TODO
func (h *Handler) UpdateInfo(ctx context.Context, req *helloworld.Info, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do UpdateInfo", acc.Name)
	}

	if len(req.Name) == 0 {
		return errors.BadRequest("Name required")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	info := model.Info{}
	err = copier.Copy(&info, req)
	if err != nil {
		logger.Errorf(ctx, "Copy.req failed %v", err)
		return errors.InternalServerError("Copy.req failed %v", err)
	}
	err = h.UpdateInfoDB(ctx, session, &info)
	if err != nil {
		logger.Errorf(ctx, "UpdateInfoDB failed %v", err)
		return errors.InternalServerError("UpdateInfoDB failed %v", err)
	}

	err = copier.Copy(rsp, &info)
	if err != nil {
		logger.Errorf(ctx, "Copy.info failed %v", err)
		return errors.InternalServerError("Copy.info failed %v", err)
	}
	return nil
}

// InsertInfo defined TODO
func (h *Handler) InsertInfo(ctx context.Context, req *helloworld.Info, rsp *helloworld.Info) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do UpdateInfo", acc.Name)
	}

	if len(req.Name) == 0 {
		return errors.BadRequest("Name required")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("InitDb failed %v", err)
	}

	info := model.Info{}
	err = copier.Copy(&info, req)
	if err != nil {
		logger.Errorf(ctx, "Copy.req failed %v", err)
		return errors.InternalServerError("Copy.req failed %v", err.Error())
	}

	err = h.InsertInfoDB(ctx, session, &info)
	if err != nil {
		logger.Errorf(ctx, "InsertSchedulePositionDB failed %v", err)
		return errors.InternalServerError("InsertSchedulePositionDB failed %v", err.Error())
	}
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
