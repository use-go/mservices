package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"

	user "proto/user"
	"user-service/handler"
)

func main() {
	// Create service
	srv := service.New(service.Name("user"))

	// Create handler
	hdl := handler.Handler{}

	// Register handler
	user.RegisterAccountHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
