package main

import (
	"assert-service/handler"
	"comm/define"
	"comm/logger"
	"comm/service/web"
)

func main() {
	hdl := handler.Handler{}
	srv := web.New(web.Name("assert"))
	srv.HandleFunc("/file/upload", hdl.FileUpload)
	srv.HandleFunc("/file/download", hdl.FileDownload)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
