package main

import (
	"comm/logger"
	"comm/service"
	"helloworld-service/handler"
	"proto/helloworld"
)

func main() {
	hdl := handler.Handler{}
	srv := service.New(service.Name("helloworld"))
	helloworld.RegisterHelloworldServiceHandler(srv.Server(), &hdl)
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
