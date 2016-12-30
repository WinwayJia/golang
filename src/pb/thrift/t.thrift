namespace go rpc
namespace java demo.rpc

service RpcService {
	// 发起远程调用
	string TwoWayCall(1:string input)
}
