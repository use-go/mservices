package main

import (
	"comm/api"
	"net/http"
)

func main() {
	srv := api.New(api.Name("test44"))
	srv.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(`{"id": 10}`))
	})
	srv.Run()
}
