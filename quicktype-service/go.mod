module quicktype-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)

replace (
	comm => ../comm
	proto => ../proto
)

replace github.com/2637309949/micro/v3 => ../../micro
