module helloworld-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	github.com/micro/micro/v3 => github.com/2637309949/micro/v3 v3.8.0
	proto => ../proto
)
