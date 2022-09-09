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
	github.com/2637309949/micro/v3 => ../../micro
	proto => ../proto
)

// fix error code PROTOCOL_ERROR
replace google.golang.org/grpc => google.golang.org/grpc v1.40.0
