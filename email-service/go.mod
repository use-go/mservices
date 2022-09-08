module email-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/Teamwork/spamc v0.0.0-20200109085853-a4e0c5c3f7a0
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/teamwork/test v0.0.0-20200108114543-02621bae84ad // indirect
	github.com/teamwork/utils v0.0.0-20211112162623-194b7eff720f // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	github.com/2637309949/micro/v3 => ../../micro
	proto => ../proto
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
