module quicktype-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	github.com/xormplus/xorm v0.0.0-20210822100304-4e1d4fcc1e67
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
)

replace (
	comm => ../comm
	proto => ../proto
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
