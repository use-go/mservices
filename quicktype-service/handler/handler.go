package handler

import (
	"fmt"
	"net/http"
	"quicktype-service/api"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	dbs map[string]*gorm.DB
}

func (h *Handler) DB(r *http.Request) (*gorm.DB, error) {
	dsn := r.URL.Query().Get("dsn")
	if len(dsn) == 0 {
		uri := (new(api.Uri)).Unmarshal(r)
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", uri.User, uri.PassWd, uri.Host, uri.Port, uri.DB)
	}
	if h.dbs[dsn] == nil {
		db, err := gorm.Open((mysql.Open(dsn)), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		h.dbs[dsn] = db
	}
	return h.dbs[dsn], nil
}
