module subscribe-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.27.1
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)
