package user_info;

message userInfo 
{
	required uint64 uid = 1;
	required string name = 2;
	optional string nickname = 3;
}
