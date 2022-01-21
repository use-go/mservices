package handler

import (
	"fmt"
	"net/http"
	"quicktype-service/api"

	"github.com/xormplus/xorm"
)

var dbs = map[string]*xorm.Engine{}

func (h *Handler) DB(r *http.Request) (*xorm.Engine, error) {
	dsn := r.URL.Query().Get("dsn")
	if len(dsn) == 0 {
		uri := (new(api.Uri)).Unmarshal(r)
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", uri.User, uri.PassWd, uri.Host, uri.Port, uri.DB)
	}
	if dbs[dsn] == nil {
		db, err := xorm.NewEngine("mysql", dsn)
		if err != nil {
			return nil, err
		}
		dbs[dsn] = db
	}
	return dbs[dsn], nil
}
