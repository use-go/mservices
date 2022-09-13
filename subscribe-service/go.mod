module subscribe-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	github.com/2637309949/micro/v3 => ../../micro
	proto => ../proto
)
