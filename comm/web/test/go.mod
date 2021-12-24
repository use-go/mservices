module test

go 1.15

require (
	github.com/2637309949/micro/v3 v3.8.2
	github.com/fatih/camelcase v1.0.0
	github.com/gorilla/mux v1.7.3
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	comm v0.0.0-00010101000000-000000000000
)

replace comm => ../../

replace github.com/2637309949/micro/v3 => ../../../../micro
