package handler

import (
	"comm/auth"
	"comm/logger"
	"comm/mark"
	whttp "comm/service/web/http"
	"fmt"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// FileUpload defined TODO
func (h *Handler) FileUpload(rw http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "FileUpload")()

	acc, ok := auth.FromContext(r.Context())
	timemark.Mark("FromContext")
	if ok {
		logger.Infof(r.Context(), "%v Do FileUpload", acc.Name)
	}
	timemark.Mark("Success")
	whttp.Success(rw, r, map[string]interface{}{"name": "hello 2012"})
}

// FileDownload defined TODO
func (h *Handler) FileDownload(rw http.ResponseWriter, r *http.Request) {
	var err error
	var timemark mark.TimeMark
	defer timemark.Init(r.Context(), "FileDownload")()

	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do FileDownload", acc.Name)
	}
	excel := excelize.NewFile()
	rw.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=\"%s-%s.xlsx\"", "newfile", time.Now().Format("2006-01-02"))
	rw.Header().Set("Content-Disposition", disposition)
	err = excel.Write(rw)
	timemark.Mark("Write")
	if err != nil {
		whttp.Fail(rw, r, err)
		return
	}
}
