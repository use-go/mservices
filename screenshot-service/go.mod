module screenshot-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.1.2
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)
