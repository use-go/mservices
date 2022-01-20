package main

import (
	"comm/define"
	"comm/logger"
	"comm/service/web"
	"quicktype-service/handler"
)

func main() {
	hdl := handler.Handler{}
	srv := web.New(web.Name("quicktype"))
	srv.HandleFunc("/type/index", hdl.Type)
	srv.HandleFunc("/type/tables", hdl.Tables)
	srv.HandleFunc("/type/table2go", hdl.Table2Go)
	srv.HandleFunc("/type/table2proto", hdl.Table2Proto)
	srv.HandleFunc("/type/table2handler", hdl.Table2Handler)
	if err := srv.Run(); err != nil {
		logger.Fatal(define.TODO, err)
	}
}
