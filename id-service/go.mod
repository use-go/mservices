module id-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.1.2
	github.com/mattheath/base62 v0.0.0-20150408093626-b80cdc656a7a // indirect
	github.com/mattheath/kala v0.0.0-20171219141654-d6276794bf0e
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
