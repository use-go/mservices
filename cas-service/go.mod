module cas-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/2637309949/micro/v3 v3.8.4
	github.com/go-oauth2/oauth2/v4 v4.4.2
	github.com/go-session/session/v3 v3.1.5
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/securecookie v1.1.1
	github.com/tidwall/buntdb v1.2.9 // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20220105145211-5b0dc2dfae98 // indirect
	google.golang.org/genproto v0.0.0-20200825200019-8632dd797987 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace (
	comm => ../comm
	proto => ../proto
)

replace github.com/2637309949/micro/v3 => ../../micro
