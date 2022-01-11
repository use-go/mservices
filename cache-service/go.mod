module cache-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d
	github.com/gin-contrib/cache v1.1.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/memcachier/mc/v3 v3.0.3
	github.com/robfig/go-cache v0.0.0-20130306151617-9fc39e0dbf62
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)
