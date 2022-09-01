package handler

import (
	"comm/errors"
	"comm/logger"
	"comm/service"
	"context"
	"user-service/model"

	"gorm.io/gorm"
)

// QueryUserDB defined TODO
func (h *Handler) QueryUserDB(ctx context.Context, session *gorm.DB, where *model.User, list *[]model.User, count ...*int64) error {
	session = session.Table(where.TableName()).Where(where).Find(list)
	if len(count) > 0 {
		session = session.Offset(0).Count(count[0])
	}
	if err := session.Error; err != nil {
		logger.Errorf(ctx, "QueryUserDB failed. [%v]", err)
		return errors.InternalServerError(service.GetName(), "QueryUserDB fail. [%v]", err)
	}
	return nil
}

// QueryUserDetailDB defined TODO
func (h *Handler) QueryUserDetailDB(ctx context.Context, session *gorm.DB, where *model.User, data *model.User) error {
	var lst []model.User
	err := h.QueryUserDB(ctx, session, where, &lst)
	if err != nil {
		logger.Errorf(ctx, "QueryUserDetailDB failed. [%s]", err.Error())
		return err
	}
	if len(lst) == 0 {
		logger.Warn(ctx, "QueryUserDetailDB empty.")
		return errors.RecordNotFound(service.GetName(), "QueryUserDetailDB empty.")
	}
	*data = lst[0]
	return nil
}

// InsertUserDB defined TODO
func (h *Handler) InsertUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	err := session.Create(data).Error
	if err != nil {
		logger.Errorf(ctx, "InsertUserDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "InsertUserDB fail. [%v]", err)
	}
	return nil
}

// UpdateUserDB defined TODO
func (h *Handler) UpdateUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	err := session.Table(data.TableName()).Model(&data).Updates(&data).Error
	if err != nil {
		logger.Errorf(ctx, "UpdateUserDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "UpdateUserDB fail. [%v]", err)
	}
	return nil
}

// SaveUserDB defined TODO
func (h *Handler) SaveUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	err := session.Save(data).Error
	if err != nil {
		logger.Errorf(ctx, "SaveUserDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "SaveUserDB fail. [%v]", err)
	}
	return nil
}

// DeleteUserDB defined TODO
func (h *Handler) DeleteUserDB(ctx context.Context, session *gorm.DB, data *model.User) error {
	err := session.Where(data).Delete(&data).Error
	if err != nil {
		logger.Errorf(ctx, "DeleteUserDB failed. [%s]", err.Error())
		return errors.InternalServerError(service.GetName(), "DeleteUserDB fail. [%v]", err)
	}
	return nil
}
