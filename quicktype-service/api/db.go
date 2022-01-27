package api

import (
	"bytes"
	"html/template"

	"github.com/xormplus/xorm/schemas"
)

func Table2DB(t *schemas.Table) (string, error) {
	message := tables2DB(t)
	msgTemplate := `package handler

import (
	"comm/errors"
	"comm/logger"
	"context"
	"helloworld-service/model"

	"gorm.io/gorm"
)

// Query{{.Name}}DB defined TODO
func (h *Handler) Query{{.Name}}DB(ctx context.Context, session *gorm.DB, where *model.{{.Name}}, list *[]*model.{{.Name}}, count ...*int64) error {
	session = session.Table(where.TableName()).Where(where).Find(list)
	if len(count) > 0 {
		session = session.Offset(0).Count(count[0])
	}
	if err := session.Error; err != nil {
		logger.Errorf(ctx, "Query{{.Name}}DB failed. [%v]", err)
		return errors.InternalServerError("Query{{.Name}}DB fail. [%v]", err)
	}
	return nil
}

// Query{{.Name}}DetailDB defined TODO
func (h *Handler) Query{{.Name}}DetailDB(ctx context.Context, session *gorm.DB, where *model.{{.Name}}, data *model.{{.Name}}) error {
	var lst []*model.{{.Name}}
	err := h.Query{{.Name}}DB(ctx, session, where, &lst)
	if err != nil {
		logger.Errorf(ctx, "Query{{.Name}}DetailDB failed. [%s]", err.Error())
		return err
	}
	if len(lst) == 0 {
		logger.Warn(ctx, "Query{{.Name}}DetailDB empty.")
		return errors.InternalServerError("Query{{.Name}}DetailDB empty.")
	}
	*data = *lst[0]
	return nil
}

// Insert{{.Name}}DB defined TODO
func (h *Handler) Insert{{.Name}}DB(ctx context.Context, session *gorm.DB, data *model.{{.Name}}) error {
	err := session.Create(data).Error
	if err != nil {
		logger.Errorf(ctx, "Insert{{.Name}}DB failed. [%s]", err.Error())
		return errors.InternalServerError("Insert{{.Name}}DB fail. [%v]", err)
	}
	return nil
}

// Update{{.Name}}DB defined TODO
func (h *Handler) Update{{.Name}}DB(ctx context.Context, session *gorm.DB, data *model.{{.Name}}) error {
	err := session.Table(data.TableName()).Model(&data).Updates(&data).Error
	if err != nil {
		logger.Errorf(ctx, "Update{{.Name}}DB failed. [%s]", err.Error())
		return errors.InternalServerError("Update{{.Name}}DB fail. [%v]", err)
	}
	return nil
}

// Save{{.Name}}DB defined TODO
func (h *Handler) Save{{.Name}}DB(ctx context.Context, session *gorm.DB, data *model.{{.Name}}) error {
	err := session.Save(data).Error
	if err != nil {
		logger.Errorf(ctx, "Save{{.Name}}DB failed. [%s]", err.Error())
		return errors.InternalServerError("Save{{.Name}}DB fail. [%v]", err)
	}
	return nil
}

// Delete{{.Name}}DB defined TODO
func (h *Handler) Delete{{.Name}}DB(ctx context.Context, session *gorm.DB, data *model.{{.Name}}) error {
	err := session.Where(data).Delete(&data).Error
	if err != nil {
		logger.Errorf(ctx, "Delete{{.Name}}DB failed. [%s]", err.Error())
		return errors.InternalServerError("Delete{{.Name}}DB fail. [%v]", err)
	}
	return nil
}
`
	tmpl, err := template.New("rw").Parse(msgTemplate)
	if err != nil {
		return "", err
	}
	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, message)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

func tables2DB(t *schemas.Table) *message {
	msg := &message{
		Name:   case2camel(t.Name),
		Fields: []*field{},
	}
	lc := len(t.Columns())
	for i := 0; i < lc; i++ {
		f := t.Columns()[i]
		newField := &field{
			Name:       toProtoFieldName(f.Name),
			TypeName:   toProtoFieldTypeNameBySql(f.SQLType),
			IsRepeated: false,
			Order:      i + 1,
		}
		msg.Fields = append(msg.Fields, newField)
	}
	return msg
}
