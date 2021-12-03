module helloworld-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.16
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)

replace github.com/micro/micro/v3 => ../comm/micro
