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

// replace github.com/2637309949/micro/v3 => ../../micro

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
// replace google.golang.org/grpc => google.golang.org/grpc v1.40.0
