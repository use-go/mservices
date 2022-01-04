package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"
	"helloworld-service/handler"
	"proto/helloworld"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hdl := handler.Handler{}
	srv := service.New(service.Name("helloworld"))
	helloworld.RegisterHelloworldHandler(srv.Server(), &hdl)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
