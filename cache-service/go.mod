module cache-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d
	github.com/go-redis/redis/v8 v8.11.4
	github.com/memcachier/mc/v3 v3.0.3
	github.com/robfig/go-cache v0.0.0-20130306151617-9fc39e0dbf62
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
