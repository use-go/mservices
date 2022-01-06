package main

import (
	"comm/logger"
	"comm/service"
	"comm/define"

	"subscribe-service/handler"
	subscribe "proto/subscribe"
)

func main() {
	// Create service
	srv := service.New(service.Name("subscribe"))

	// Create handler
	hdl := handler.Handler{}

	// Register handler
	subscribe.RegisterSubscribeHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
