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
func (h *Handler) DeleteInfo(ctx context.Context, req *helloworld.DeleteInfoRequest, rsp *helloworld.DeleteInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteInfo", acc.Name)
	}

	if req.Id == 0 {
		return errors.BadRequest("Id required")
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	where := model.Info{
		Id: req.Id,
	}
	err = h.DeleteInfoDB(ctx, session, &where)
	if err != nil {
		logger.Errorf(ctx, "DeleteInfoDB failed. [%v]", err)
		return errors.InternalServerError("deleteInfoDB failed %v", err.Error())
	}
	rsp.Id = where.Id
	return nil
}

// UpdateInfo defined TODO
func (h *Handler) UpdateInfo(ctx context.Context, req *helloworld.UpdateInfoRequest, rsp *helloworld.UpdateInfoResponse) error {
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
func (h *Handler) InsertInfo(ctx context.Context, req *helloworld.InsertInfoRequest, rsp *helloworld.InsertInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do InsertInfo", acc.Name)
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
func (h *Handler) QueryInfoDetail(ctx context.Context, req *helloworld.QueryInfoDetailRequest, rsp *helloworld.QueryInfoDetailResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfoDetail", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("InitDb failed %v", err)
	}

	where := model.Info{
		Id: req.Id,
	}
	info := model.Info{}
	err = h.QueryInfoDetailDB(ctx, session, &where, &info)
	if err != nil {
		return errors.InternalServerError("QueryInfoDetailDB failed %v", err)
	}
	err = copier.Copy(rsp, &info)
	if err != nil {
		return errors.InternalServerError("copier.Copy failed %v", err)
	}
	return nil
}

// QueryInfo defined TODO
func (h *Handler) QueryInfo(ctx context.Context, req *helloworld.QueryInfoRequest, rsp *helloworld.QueryInfoResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfo", acc.Name)
	}

	session, err := db.InitDb(ctx)
	session = db.SetLimit(ctx, session, req)
	if err != nil {
		return errors.InternalServerError("InitDb failed %v", err)
	}

	var totalCount int64
	var lst []*model.Info
	where := model.Info{
		Name: req.GetName(),
	}
	err = h.QueryInfoDB(ctx, session, &where, &lst, &totalCount)
	if err != nil {
		return errors.InternalServerError("QueryInfoDB failed %v", err)
	}
	err = copier.Copy(&rsp.Data, &lst)
	if err != nil {
		return errors.InternalServerError("copier.Copy failed %v", err)
	}

	rsp.TotalCount = totalCount
	rsp.Page = totalCount / req.Size
	if totalCount%req.Size != 0 {
		rsp.Page += 1
	}
	return nil
}
