package main

import (
	"comm/logger"
	"comm/service"
	"comm/define"

	"chat-service/handler"
	chat "proto/chat"
)

func main() {
	// Create service
	srv := service.New(service.Name("chat"))

	// Create handler
	hdl := handler.Handler{}

	// Register handler
	chat.RegisterChatHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
