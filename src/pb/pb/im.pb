package im;

enum XMessageType {
	Login = 0;
	Logout = 1;
	Ping = 2;
	Pong = 3;
	Chat = 4;
};

message  XMessageHeader {
	optional XMessageType type = 1;
	optional string token = 2;
	optional int32 length = 3; // Body length, not include the header
};

message XLoginReq {
	optional int64 UserID = 1;
};

message XLoginResp {
	optional string token = 1;
	optional base.ErrInfo errmsg = 2;
};

service IMService {
	rpc Login(XLoginReq) returns (XLoginResp);
}
