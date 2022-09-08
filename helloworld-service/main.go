package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"
	"comm/store"
	"helloworld-service/handler"
	"proto/cache"
	"proto/email"
	"proto/helloworld"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	srv := service.New(service.Name("helloworld"))
	hdl := handler.Handler{
		CacheService: store.CacheService(cache.NewCacheService("cache", srv.Client())),
		EmailService: email.NewEmailService("email", srv.Client()),
	}
	helloworld.RegisterHelloworldHandler(srv.Server(), &hdl)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
