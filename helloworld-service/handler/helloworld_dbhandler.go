package handler

import (
	"comm/errors"
	"comm/logger"
	"context"
	"helloworld-service/model"

	"github.com/jinzhu/gorm"
)

// QueryInfoDB defined TODO
func (s *Handler) QueryInfoDB(ctx context.Context, session *gorm.DB, where *model.Info, list *[]model.Info, count ...*int32) error {
	session = session.Table(where.TableName()).Where(where).Find(list)
	if len(count) > 0 {
		session = session.Offset(0).Count(count[0])
	}
	if errs := session.GetErrors(); len(errs) != 0 {
		logger.Errorf(ctx, "QueryInfoDB failed. [%v]", errs)
		return errors.InternalServerError("QueryInfoDB fail. [%v]", errs)
	}
	return nil
}

// QueryInfoDetailDB defined TODO
func (s *Handler) QueryInfoDetailDB(ctx context.Context, session *gorm.DB, where *model.Info, data *model.Info) error {
	var lst []model.Info
	err := s.QueryInfoDB(ctx, session, where, &lst)
	if err != nil {
		logger.Errorf(ctx, "QueryInfoDetailDB failed. [%s]", err.Error())
		return err
	}
	if len(lst) == 0 {
		logger.Warn(ctx, "QueryInfoDetailDB empty.")
		return errors.InternalServerError("QueryInfoDetailDB empty.")
	}
	*data = lst[0]
	return nil
}

// InsertInfoDB defined TODO
func (s *Handler) InsertInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Create(data).Error
	if err != nil {
		logger.Errorf(ctx, "InsertInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError("InsertInfoDB fail. [%v]", err)
	}
	return nil
}

// UpdateInfoDB defined TODO
func (s *Handler) UpdateInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Table(data.TableName()).Model(&data).Updates(&data).Error
	if err != nil {
		logger.Errorf(ctx, "UpdateInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError("UpdateInfoDB fail. [%v]", err)
	}
	return nil
}

// SaveInfoDB defined TODO
func (s *Handler) SaveInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Save(data).Error
	if err != nil {
		logger.Errorf(ctx, "SaveInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError("SaveInfoDB fail. [%v]", err)
	}
	return nil
}

// DeleteInfoDB defined TODO
func (s *Handler) DeleteInfoDB(ctx context.Context, session *gorm.DB, data *model.Info) error {
	err := session.Where(data).Delete(&data).Error
	if err != nil {
		logger.Errorf(ctx, "DeleteInfoDB failed. [%s]", err.Error())
		return errors.InternalServerError("DeleteInfoDB fail. [%v]", err)
	}
	return nil
}
