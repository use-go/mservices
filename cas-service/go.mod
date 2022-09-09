module cas-service

go 1.15

require (
	comm v0.0.0-00010101000000-000000000000
	github.com/2637309949/micro/v3 v3.8.9
	github.com/go-oauth2/oauth2/v4 v4.4.2
	github.com/go-session/session/v3 v3.1.5
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gorilla/securecookie v1.1.1
	github.com/tidwall/buntdb v1.2.9 // indirect
)

replace (
	comm => ../comm
	github.com/2637309949/micro/v3 => ../../micro
	proto => ../proto
)

// fix error code PROTOCOL_ERROR
replace google.golang.org/grpc => google.golang.org/grpc v1.40.0
