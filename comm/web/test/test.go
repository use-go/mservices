package main

import (
	"comm/web"
	"net/http"
)

func main() {
	srv := web.New(web.Name("test44"))
	srv.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(`{"id": 10}`))
	})
	srv.Run()
}
