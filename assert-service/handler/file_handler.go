package handler

import (
	"comm/auth"
	"comm/logger"
	whttp "comm/service/web/http"
	"fmt"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// FileUpload defined TODO
func (h *Handler) FileUpload(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do FileUpload", acc.Name)
	}
	whttp.WriteJSON(rw, r, map[string]interface{}{"name": "hello 2012"})
}

// FileDownload defined TODO
func (h *Handler) FileDownload(rw http.ResponseWriter, r *http.Request) {
	acc, ok := auth.FromContext(r.Context())
	if ok {
		logger.Infof(r.Context(), "%v Do FileDownload", acc.Name)
	}
	excel := excelize.NewFile()
	rw.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=\"%s-%s.xlsx\"", "newfile", time.Now().Format("2006-01-02"))
	rw.Header().Set("Content-Disposition", disposition)
	err := excel.Write(rw)
	if err != nil {
		whttp.WriteError(rw, r, err)
		return
	}
}
