package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"
	"comm/store"

	"proto/cache"
	"proto/email"
	user "proto/user"
	"user-service/handler"
)

func main() {
	// Create service
	srv := service.New(service.Name("user"))

	// Create handler
	hdl := handler.Handler{
		CacheService: store.CacheService(cache.NewCacheService("cache", srv.Client())),
		EmailService: email.NewEmailService("email", srv.Client()),
	}

	// Register handler
	user.RegisterAccountHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
