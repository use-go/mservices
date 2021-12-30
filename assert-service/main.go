package main

import (
	"assert-service/handler"
	"comm/logger"
	"comm/web"
)

func main() {
	hdl := handler.Handler{}
	srv := web.New(web.Name("assert"))
	srv.HandleFunc("/file/upload", hdl.FileUpload)
	srv.HandleFunc("/file/download", hdl.FileDownload)
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
