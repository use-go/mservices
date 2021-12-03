package template

var (
	ProtoServiceSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

import "proto/{{dehyphen .Alias}}/{{dehyphen .Alias}}.proto";

service {{title .Alias}}Service {
	rpc Call(Request) returns (Response) {}
}
`
	ProtoModelSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
`
)
