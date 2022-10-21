package handler

import (
	"comm/auth"
	"comm/logger"
	"comm/mark"
	whttp "comm/service/web/http"
	"errors"
	"net/http"
	"quicktype-service/model/api"
)

// Index defined TODO
func (h *Handler) Index(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Index")()

	acc, ok := auth.FromContext(r.Context())
	timemark.Mark("FromContext")
	if ok {
		logger.Infof(r.Context(), "%v Do Index", acc.Name)
	}

	whttp.OutputHTML(rw, r, "static/index.html")
}

// Tables defined TODO
func (h *Handler) Tables(rw http.ResponseWriter, r *http.Request) {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Tables")()

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
	timemark.Mark("DBMetas")
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
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Table2Go")()

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
	timemark.Mark("DBMetas")
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
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Table2Proto")()

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
	timemark.Mark("DBMetas")
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
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Table2Handler")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2Handler", acc.Name)
	}
	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	table := r.URL.Query().Get("table")
	schemas, err := db.DBMetas()
	timemark.Mark("DBMetas")
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
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Table2RW")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2RW", acc.Name)
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

// Table2DB defined TODO
func (h *Handler) Table2DB(rw http.ResponseWriter, r *http.Request) {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "Table2DB")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do Table2DB", acc.Name)
	}
	db, err := h.DB(r)
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	table := r.URL.Query().Get("table")
	schemas, err := db.DBMetas()
	timemark.Mark("DBMetas")
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
	for _, schema := range schemas {
		if schema.Name == table {
			proto, err := api.Table2DB(schema)
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

// Logout defined TODO
func (h *Handler) Logout(rw http.ResponseWriter, r *http.Request) {
}
