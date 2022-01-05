module email-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/Teamwork/spamc v0.0.0-20200109085853-a4e0c5c3f7a0
	github.com/teamwork/test v0.0.0-20200108114543-02621bae84ad // indirect
	github.com/teamwork/utils v0.0.0-20211112162623-194b7eff720f // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	proto v0.0.0-00010101000000-000000000000
)

replace (
	comm => ../comm
	proto => ../proto
)

replace github.com/2637309949/micro/v3 => ../../micro
