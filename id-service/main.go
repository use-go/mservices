package main

import (
	"comm/define"
	"comm/logger"
	"comm/service"
	"id-service/handler"
	"math/rand"
	"proto/id"

	"github.com/mattheath/kala/bigflake"
	"github.com/mattheath/kala/snowflake"
)

func main() {
	// Create service
	srv := service.New(service.Name("id"))

	// Create handler
	it := rand.Intn(100)
	sf, err := snowflake.New(uint32(it))
	if err != nil {
		logger.Fatal(define.TODO, err)
	}
	bg, err := bigflake.New(uint64(it))
	if err != nil {
		logger.Fatal(define.TODO, err)
	}

	hdl := handler.Handler{Snowflake: sf, Bigflake: bg}

	// Register handler
	id.RegisterIdHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
