module helloworld-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/copier v0.3.4
	gorm.io/gorm v1.22.5
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)

replace github.com/2637309949/micro/v3 => ../../micro
