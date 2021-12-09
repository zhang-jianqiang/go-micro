package template

// ProtoFNC is the .proto file template used for new function projects.
var ProtoFNC = `syntax = "proto3";

package {{dehyphen .Service}};

option go_package = "./proto;{{dehyphen .Service}}";

service {{title .Service}} {
	rpc Call(CallRequest) returns (CallResponse) {}
}

message CallRequest {
	string name = 1;
}

message CallResponse {
	string msg = 1;
}
`

// ProtoSRV is the .proto file template used for new service projects.
var ProtoSRV = `syntax = "proto3";

package {{dehyphen .Service}};

option go_package = "./;{{dehyphen .Service}}";

service {{title .Service}} {
	rpc FindOneById(FindOneByIdRequest) returns (FindOneByIdRequestResponse) {}
}

message FindOneByIdRequest {
	int32 id = 1;
}

message FindOneByIdRequestResponse {
	int32 id = 1;
}
`
