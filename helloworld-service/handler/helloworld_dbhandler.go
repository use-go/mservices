package handler

import (
	"comm/errors"
	"comm/logger"
	"context"
	"helloworld-service/model"

	"github.com/jinzhu/gorm"
)

func (s *Handler) QueryHelloworldDB(ctx context.Context, session *gorm.DB, where *model.Helloworld, list *[]model.Helloworld, count ...*int32) error {
	session = session.Table(where.TableName()).Where(where).Find(list)
	if len(count) > 0 {
		session = session.Offset(0).Count(count[0])
	}
	if errs := session.GetErrors(); len(errs) != 0 {
		logger.Errorf("QueryHelloworldDB failed. [%v]", errs)
		return errors.New(50001, "QueryHelloworldDB fail. [%v]", errs)
	}
	return nil
}

func (s *Handler) QueryHelloworldDetailDB(ctx context.Context, session *gorm.DB, where *model.Helloworld, data *model.Helloworld) error {
	var lst []model.Helloworld
	err := s.QueryHelloworldDB(ctx, session, where, &lst)
	if err != nil {
		logger.Errorf("QueryHelloworldDetailDB failed. [%s]", err.Error())
		return err
	}
	if len(lst) == 0 {
		logger.Warn("QueryHelloworldDetailDB empty.")
		return errors.New(50001, "QueryHelloworldDetailDB empty.")
	}
	*data = lst[0]
	return nil
}

func (s *Handler) InsertHelloworldDB(ctx context.Context, session *gorm.DB, data *model.Helloworld) error {
	err := session.Create(data).Error
	if err != nil {
		logger.Errorf("InsertHelloworldDB failed. [%s]", err.Error())
		return errors.New(50001, "InsertHelloworldDB failed. [%s]", err.Error())
	}
	return nil
}

func (s *Handler) UpdateHelloworldDB(ctx context.Context, session *gorm.DB, data *model.Helloworld) error {
	err := session.Table(data.TableName()).Model(&data).Updates(&data).Error
	if err != nil {
		logger.Errorf("UpdateHelloworldDB failed. [%s]", err.Error())
		return errors.New(50001, "UpdateHelloworldDB failed. [%s]", err.Error())
	}
	return nil
}

func (s *Handler) SaveHelloworldDB(ctx context.Context, session *gorm.DB, data *model.Helloworld) error {
	err := session.Save(data).Error
	if err != nil {
		logger.Errorf("SaveHelloworldDB failed. [%s]", err.Error())
		return errors.New(50001, "SaveHelloworldDB failed. [%s]", err.Error())
	}
	return nil
}

func (s *Handler) DeleteHelloworldDB(ctx context.Context, session *gorm.DB, data *model.Helloworld) error {
	err := session.Where(data).Delete(&data).Error
	if err != nil {
		logger.Errorf("DeleteHelloworldDB failed. [%s]", err.Error())
		return errors.New(50001, "DeleteHelloworldDB failed. [%s]", err.Error())
	}
	return nil
}
