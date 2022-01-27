package handler

import (
	"comm/auth"
	"comm/logger"
	whttp "comm/service/web/http"
	"errors"
	"net/http"
	"quicktype-service/api"
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
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Tables", acc.Name)
	}

	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	schemas, err := db.DBMetas()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var tables []string
	for _, schema := range schemas {
		tables = append(tables, schema.Name)
	}

	whttp.Success(rw, r, tables)
}

// Table2Go defined TODO
func (h *Handler) Table2Go(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Go", acc.Name)
	}
	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	table := r.URL.Query().Get("table")
	schemas, err := db.DBMetas()
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	for _, schema := range schemas {
		if schema.Name == table {
			proto, err := api.Table2Struct(schema)
			if err != nil {
				whttp.Fail(rw, r, err)
				return
			}
			whttp.Success(rw, r, proto)
			return
		}
	}
	whttp.Fail(rw, r, errors.New("not found"))
}

// Table2Proto defined TODO
func (h *Handler) Table2Proto(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Proto", acc.Name)
	}
	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	table := r.URL.Query().Get("table")
	schemas, err := db.DBMetas()
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	for _, schema := range schemas {
		if schema.Name == table {
			proto, err := api.Table2Proto(schema)
			if err != nil {
				whttp.Fail(rw, r, err)
				return
			}
			whttp.Success(rw, r, proto)
			return
		}
	}
	whttp.Fail(rw, r, errors.New("not found"))
}

// Table2Handler defined TODO
func (h *Handler) Table2Handler(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Proto", acc.Name)
	}
	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	table := r.URL.Query().Get("table")
	schemas, err := db.DBMetas()
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	for _, schema := range schemas {
		if schema.Name == table {
			proto, err := api.Table2Handler(schema)
			if err != nil {
				whttp.Fail(rw, r, err)
				return
			}
			whttp.Success(rw, r, proto)
			return
		}
	}
	whttp.Fail(rw, r, errors.New("not found"))
}

// Table2RW defined TODO
func (h *Handler) Table2RW(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Proto", acc.Name)
	}
	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	table := r.URL.Query().Get("table")
	schemas, err := db.DBMetas()
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	for _, schema := range schemas {
		if schema.Name == table {
			proto, err := api.Table2RW(schema)
			if err != nil {
				whttp.Fail(rw, r, err)
				return
			}
			whttp.Success(rw, r, proto)
			return
		}
	}
	whttp.Fail(rw, r, errors.New("not found"))
}
