package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"
	"time"

	"cache-service/api"
	"cache-service/handler"
	"proto/cache"
)

func main() {
	// Create service
	srv := service.New(service.Name("cache"))

	// Create handler
	hdl := handler.Handler{
		CacheStore: api.NewInMemoryStore(60 * time.Second),
	}

	// Register handler
	cache.RegisterCacheHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
