package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"

	"proto/id"
	"proto/screenshot"
	"screenshot-service/handler"
)

func main() {
	// Create service
	srv := service.New(service.Name("screenshot"))

	// Create handler
	hdl := handler.Handler{
		IdService: id.NewIdService("id", srv.Client()),
	}

	// Register handler
	screenshot.RegisterScreenshotHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
