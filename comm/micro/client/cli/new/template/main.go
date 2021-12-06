package template

var (
	MainSRV = `package main

import (
	"comm/logger"
	"comm/service"

	"{{.Dir}}-service/handler"
	{{dehyphen .Alias}} "proto/{{.Dir}}"
)

func main() {
	// Create handler
	hdl := handler.Handler{}

	// Create service
	srv := service.New(service.Name("{{lower .Alias}}"))

	// Register handler
	{{dehyphen .Alias}}.Register{{title .Alias}}ServiceHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
`
)
