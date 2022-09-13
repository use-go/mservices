module assert-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
)

replace (
	comm => ../comm
	github.com/2637309949/micro/v3 => ../../micro
	proto => ../proto
)
