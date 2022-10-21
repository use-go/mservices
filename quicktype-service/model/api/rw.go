package api

import (
	"bytes"
	"html/template"

	"github.com/xormplus/xorm/schemas"
)

func Table2RW(t *schemas.Table) (string, error) {
	message := tables2RW(t)
	msgTemplate := `package handler

import (
	"comm/auth"
	"comm/db"
	"comm/errors"
	"comm/logger"
	"comm/util"
	"context"
	"xxx-service/model"
	"proto/xxx"
	"time"
)

// Delete{{.Name}} defined todo
func (h *Handler) Delete{{.Name}}(ctx context.Context, req *xxx.Delete{{.Name}}Request, rsp *xxx.Delete{{.Name}}Response) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Delete{{.Name}}", acc.Name)
	}

	err := util.IsZero(req, "id")
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb failed. [%v]", err)
		return err
	}

	where := model.{{.Name}}{
		Id: req.Id,
	}
	err = h.Delete{{.Name}}DB(ctx, session, &where)
	if err != nil {
		logger.Errorf(ctx, "Delete{{.Name}}DB failed. [%v]", err)
		return errors.InternalServerError("delete{{.Name}}DB failed %v", err.Error())
	}
	rsp.Id = where.Id
	return nil
}

// Update{{.Name}} defined todo
func (h *Handler) Update{{.Name}}(ctx context.Context, req *xxx.Update{{.Name}}Request, rsp *xxx.Update{{.Name}}Response) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Update{{.Name}}", acc.Name)
	}

	err := util.IsZero(req, "id")
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb failed. [%v]", err)
		return err
	}

	info := model.{{.Name}}{}
	err = info.Unmarshal(req)
	if err != nil {
		logger.Errorf(ctx, "Unmarshal failed %v", err)
		return errors.InternalServerError("Unmarshal failed %v", err)
	}
	err = h.Update{{.Name}}DB(ctx, session, &info)
	if err != nil {
		logger.Errorf(ctx, "Update{{.Name}}DB failed %v", err)
		return errors.InternalServerError("Update{{.Name}}DB failed %v", err)
	}

	err = info.Marshal(rsp)
	if err != nil {
		logger.Errorf(ctx, "Marshal failed %v", err)
		return errors.InternalServerError("Marshal failed %v", err)
	}
	return nil
}

// Insert{{.Name}} defined todo
func (h *Handler) Insert{{.Name}}(ctx context.Context, req *xxx.Insert{{.Name}}Request, rsp *xxx.Insert{{.Name}}Response) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Insert{{.Name}}", acc.Name)
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb failed. [%v]", err)
		return err
	}

	info := model.{{.Name}}{}
	err = info.Unmarshal(req)
	if err != nil {
		logger.Errorf(ctx, "Unmarshal failed %v", err)
		return errors.InternalServerError("Unmarshal failed %v", err.Error())
	}

	info.CreatedAt = time.Now()
	info.UpdatedAt = time.Now()
	err = h.Insert{{.Name}}DB(ctx, session, &info)
	if err != nil {
		logger.Errorf(ctx, "InsertSchedulePositionDB failed %v", err)
		return errors.InternalServerError("InsertSchedulePositionDB failed %v", err.Error())
	}
	return nil
}

// Query{{.Name}}Detail defined todo
func (h *Handler) Query{{.Name}}Detail(ctx context.Context, req *xxx.Query{{.Name}}DetailRequest, rsp *xxx.Query{{.Name}}DetailResponse) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Query{{.Name}}Detail", acc.Name)
	}

	err := util.IsZero(req, "id")
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	session, err := db.InitDb(ctx)
	if err != nil {
		logger.Errorf(ctx, "InitDb failed. [%v]", err)
		return err
	}

	where := model.{{.Name}}{
		Id: req.Id,
	}
	info := model.{{.Name}}{}
	err = h.Query{{.Name}}DetailDB(ctx, session, &where, &info)
	if err != nil {
		return errors.InternalServerError("Query{{.Name}}DetailDB failed %v", err)
	}
	err = info.Marshal(rsp)
	if err != nil {
		logger.Errorf(ctx, "Marshal failed %v", err)
		return errors.InternalServerError("Marshal failed %v", err)
	}
	return nil
}

// Query{{.Name}} defined todo
func (h *Handler) Query{{.Name}}(ctx context.Context, req *xxx.Query{{.Name}}Request, rsp *xxx.Query{{.Name}}Response) error {
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Query{{.Name}}", acc.Name)
	}

	session, err := db.InitDb(ctx)
	session = db.SetLimit(ctx, session, req)
	if err != nil {
		logger.Errorf(ctx, "InitDb failed. [%v]", err)
		return err
	}

	var totalCount int64
	var lst []*model.{{.Name}}
	where := model.{{.Name}}{
		Name: req.GetName(),
	}
	err = h.Query{{.Name}}DB(ctx, session, &where, &lst, &totalCount)
	if err != nil {
		return errors.InternalServerError("Query{{.Name}}DB failed %v", err)
	}

	err = model.{{.Name}}UnmarshalLst(&lst, &rsp.Data)
	if err != nil {
		return errors.InternalServerError("{{.Name}}UnmarshalLst failed %v", err)
	}

	rsp.TotalCount = totalCount
	rsp.Page = totalCount / req.Size
	if totalCount%req.Size != 0 {
		rsp.Page += 1
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

func tables2RW(t *schemas.Table) *message {
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
