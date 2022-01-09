package main

import (
	"comm/logger"
	"comm/service"
	"comm/define"

	"user-service/handler"
	user "proto/user"
)

func main() {
	// Create service
	srv := service.New(service.Name("user"))

	// Create handler
	hdl := handler.Handler{}

	// Register handler
	user.RegisterUserHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
