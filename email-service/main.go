package main

import (
	"comm/config"
	"comm/define"
	"comm/logger"
	"comm/service"

	"email-service/handler"
	"proto/email"

	"github.com/Teamwork/spamc"
)

func main() {
	// Create server
	srv := service.New(service.Name("email"))

	sgc := handler.Smtp{}
	val, err := config.Get("smtp")
	if err != nil {
		logger.Warn(define.TODO, err)
	}
	if err := val.Scan(&sgc); err != nil {
		logger.Fatal(define.TODO, err)
	}

	scf := handler.SpamcConf{}
	val, err = config.Get("spam")
	if err != nil {
		logger.Warn(define.TODO, err)
	}

	if err := val.Scan(&scf); err != nil {
		logger.Fatal(define.TODO, err)
	}

	// Create handler
	hdl := handler.Handler{
		Smtp:  &sgc,
		Spamc: spamc.New(scf.SpamdAddress, nil),
	}
	// Registe service
	email.RegisterEmailHandler(srv.Server(), &hdl)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
