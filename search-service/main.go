package main

import (
	"comm/config"
	"comm/define"
	"comm/logger"
	"comm/service"
	"crypto/tls"
	"net/http"
	search "proto/search"
	"search-service/handler"

	open "github.com/opensearch-project/opensearch-go"
)

func main() {
	// Create service
	srv := service.New(service.Name("search"))

	// Create handler
	v, err := config.Get("search")
	if err != nil {
		logger.Fatalf(define.TODO, "Failed to load config %s", err)
	}
	var c handler.Conf
	if err := v.Scan(&c); err != nil {
		logger.Fatalf(define.TODO, "Failed to load config %s", err)
	}
	if len(c.OpenAddr) == 0 || len(c.User) == 0 || len(c.Pass) == 0 {
		logger.Fatalf(define.TODO, "Missing configuration")
	}

	oc := open.Config{
		Addresses: []string{c.OpenAddr},
		Username:  c.User,
		Password:  c.Pass,
	}
	if c.Insecure {
		oc.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For testing only. Use certificate for validation.
		}
	}

	client, err := open.NewClient(oc)
	if err != nil {
		logger.Fatalf(define.TODO, "Error configuring search client %s", err)
	}
	hdl := handler.Handler{
		Client: client,
	}

	// Register handler
	search.RegisterSearchHandler(srv.Server(), &hdl)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
