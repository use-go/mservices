package main

import (
	"comm/define"
	"comm/logger"
	"comm/service/web"
	"quicktype-service/handler"
)

func main() {
	hdl := handler.Handler{}
	srv := web.New(web.Name("quicktype"))
	srv.HandleFunc("/type/index", hdl.Type)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
