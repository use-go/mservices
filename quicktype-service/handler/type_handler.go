package handler

import (
	"comm/auth"
	"comm/logger"
	whttp "comm/service/web/http"
	"fmt"
	"net/http"
	"quicktype-service/api"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Type defined TODO
func (h *Handler) Type(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Type", acc.Name)
	}

	whttp.OutputHTML(rw, r, "static/index.html")
}

// Tables defined TODO
func (h *Handler) Tables(rw http.ResponseWriter, r *http.Request) {
	var err error
	var db *gorm.DB
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Tables", acc.Name)
	}

	dsn := r.URL.Query().Get("dsn")
	if len(dsn) > 0 {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		uri := (new(api.Uri)).Unmarshal(r)
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", uri.User, uri.PassWd, uri.Host, uri.Port, uri.DB)
		db, err = gorm.Open((mysql.Open(dsn)), &gorm.Config{})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	db = db.Debug()
	rows, err := db.Table("information_schema.tables").Select("table_name").Where("table_schema = ?", "public").Rows()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var tables []string
	var name string
	for rows.Next() {
		rows.Scan(&name)
		tables = append(tables, name)
	}
	logger.Infof(r.Context(), "tables=%v", tables)
	whttp.WriteJSON(rw, r, tables)
}

// Table2Go defined TODO
func (h *Handler) Table2Go(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Go", acc.Name)
	}

	whttp.OutputHTML(rw, r, "static/index.html")
}

// Table2Proto defined TODO
func (h *Handler) Table2Proto(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Proto", acc.Name)
	}

	whttp.OutputHTML(rw, r, "static/index.html")
}

// Table2Handler defined TODO
func (h *Handler) Table2Handler(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Handler", acc.Name)
	}

	whttp.OutputHTML(rw, r, "static/index.html")
}
