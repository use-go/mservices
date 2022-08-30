package handler

import (
	"comm/auth"
	"comm/db"
	"comm/errors"
	"comm/logger"
	"comm/mark"
	"comm/util"
	"context"
	"helloworld-service/model"
	"proto/helloworld"
	"time"
)

// DeleteInfo defined TODO
func (h *Handler) DeleteInfo(ctx context.Context, req *helloworld.DeleteInfoRequest, rsp *helloworld.DeleteInfoResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "DeleteInfo")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteInfo", acc.Name)
	}

	err = util.IsZero(req, "id")
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	session, err := db.InitDb(ctx)
	timemark.Mark("InitDb")
	if err != nil {
		return errors.InternalServerError("init db error %v", err)
	}

	where := model.Info{
		Id: req.Id,
	}
	err = h.DeleteInfoDB(ctx, session, &where)
	timemark.Mark("DeleteInfoDB")
	if err != nil {
		logger.Errorf(ctx, "DeleteInfoDB failed. [%v]", err)
		return errors.InternalServerError("deleteInfoDB failed %v", err.Error())
	}
	rsp.Id = where.Id
	return nil
}

// UpdateInfo defined TODO
func (h *Handler) UpdateInfo(ctx context.Context, req *helloworld.UpdateInfoRequest, rsp *helloworld.UpdateInfoResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "UpdateInfo")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do UpdateInfo", acc.Name)
	}

	err = util.IsZero(req, "id")
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	session, err := db.InitDb(ctx)
	timemark.Mark("InitDb")
	if err != nil {
		return err
	}

	info := model.Info{}
	err = info.Unmarshal(req)
	if err != nil {
		logger.Errorf(ctx, "Unmarshal failed %v", err)
		return errors.InternalServerError("Unmarshal failed %v", err)
	}
	err = h.UpdateInfoDB(ctx, session, &info)
	timemark.Mark("UpdateInfoDB")
	if err != nil {
		logger.Errorf(ctx, "UpdateInfoDB failed %v", err)
		return errors.InternalServerError("UpdateInfoDB failed %v", err)
	}

	err = info.Marshal(rsp)
	if err != nil {
		logger.Errorf(ctx, "Marshal failed %v", err)
		return errors.InternalServerError("Marshal failed %v", err)
	}
	return nil
}

// InsertInfo defined TODO
func (h *Handler) InsertInfo(ctx context.Context, req *helloworld.InsertInfoRequest, rsp *helloworld.InsertInfoResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "InsertInfo")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do InsertInfo", acc.Name)
	}

	session, err := db.InitDb(ctx)
	timemark.Mark("InitDb")
	if err != nil {
		return errors.InternalServerError("InitDb failed %v", err)
	}

	info := model.Info{}
	err = info.Unmarshal(req)
	if err != nil {
		logger.Errorf(ctx, "Unmarshal failed %v", err)
		return errors.InternalServerError("Unmarshal failed %v", err.Error())
	}

	info.CreatedAt = time.Now()
	info.UpdatedAt = time.Now()
	err = h.InsertInfoDB(ctx, session, &info)
	timemark.Mark("InsertInfoDB")
	if err != nil {
		logger.Errorf(ctx, "InsertSchedulePositionDB failed %v", err)
		return errors.InternalServerError("InsertSchedulePositionDB failed %v", err.Error())
	}
	return nil
}

// QueryInfoDetail defined TODO
func (h *Handler) QueryInfoDetail(ctx context.Context, req *helloworld.QueryInfoDetailRequest, rsp *helloworld.QueryInfoDetailResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "QueryInfoDetail")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfoDetail", acc.Name)
	}

	err = util.IsZero(req, "id")
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	session, err := db.InitDb(ctx)
	timemark.Mark("InitDb")
	if err != nil {
		return errors.InternalServerError("InitDb failed %v", err)
	}

	where := model.Info{
		Id: req.Id,
	}
	info := model.Info{}
	err = h.QueryInfoDetailDB(ctx, session, &where, &info)
	timemark.Mark("QueryInfoDetailDB")
	if err != nil {
		return errors.InternalServerError("QueryInfoDetailDB failed %v", err)
	}
	err = info.Marshal(rsp)
	if err != nil {
		logger.Errorf(ctx, "Marshal failed %v", err)
		return errors.InternalServerError("Marshal failed %v", err)
	}
	return nil
}

// QueryInfo defined TODO
func (h *Handler) QueryInfo(ctx context.Context, req *helloworld.QueryInfoRequest, rsp *helloworld.QueryInfoResponse) error {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(ctx, "QueryInfo")()

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do QueryInfo", acc.Name)
	}

	timemark.Mark("InitDb")
	session, err := db.InitDb(ctx)
	if err != nil {
		return errors.InternalServerError("InitDb failed %v", err)
	}
	session = db.SetLimit(ctx, session, req)
	session = db.SetOrder(ctx, session, req)

	var totalCount int64
	var lst []*model.Info
	where := model.Info{
		Name: req.GetName(),
	}
	err = h.QueryInfoDB(ctx, session, &where, &lst, &totalCount)
	timemark.Mark("QueryInfoDB")
	if err != nil {
		return errors.InternalServerError("QueryInfoDB failed %v", err)
	}

	err = model.InfoUnmarshalLst(&lst, &rsp.Data)
	if err != nil {
		return errors.InternalServerError("InfoUnmarshalLst failed %v", err)
	}

	rsp.TotalCount = totalCount
	rsp.Page = totalCount / req.Size
	if totalCount%req.Size != 0 {
		rsp.Page += 1
	}
	return nil
}
