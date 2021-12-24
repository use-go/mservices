package main

import (
	"comm/web"
	"net/http"
)

func main() {
	srv := web.New(web.Name("test"))
	srv.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
	})
	srv.Run(":9081")
}
