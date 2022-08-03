package main

import (
	"comm/logger"
	"comm/service"
	"comm/define"

	"avatar-service/handler"
	avatar "proto/avatar"
)

func main() {
	// Create service
	srv := service.New(service.Name("avatar"))

	// Create handler
	hdl := handler.Handler{}

	// Register handler
	avatar.RegisterAvatarHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
