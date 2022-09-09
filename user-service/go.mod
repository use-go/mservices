module user-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/jinzhu/copier v0.3.4
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	gorm.io/gorm v1.22.5
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	github.com/2637309949/micro/v3 => ../../micro
	proto => ../proto
)

// fix error code PROTOCOL_ERROR
replace google.golang.org/grpc => google.golang.org/grpc v1.40.0
