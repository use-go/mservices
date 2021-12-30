package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// FileUpload defined TODO
func (h *Handler) FileUpload(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(`{"id": 10}`))
}

// FileDownload defined TODO
func (h *Handler) FileDownload(rw http.ResponseWriter, r *http.Request) {
	excel := excelize.NewFile()
	rw.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=\"%s-%s.xlsx\"", "newfile", time.Now().Format("2006-01-02"))
	rw.Header().Set("Content-Disposition", disposition)
	err := excel.Write(rw)
	if err != nil {
		writeError(rw, err)
		return
	}
}
