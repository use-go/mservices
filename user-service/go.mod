module user-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/jinzhu/copier v0.3.4
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)
