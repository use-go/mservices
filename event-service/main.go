package main

import (
	"comm/logger"
	"comm/service"
	"comm/define"

	"event-service/handler"
	event "proto/event"
)

func main() {
	// Create service
	srv := service.New(service.Name("event"))

	// Create handler
	hdl := handler.Handler{}

	// Register handler
	event.RegisterEventHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
