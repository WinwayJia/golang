package main
 
import (
    "cltx/rpc"
	"cltx/UserRmbRpcDef"
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/golang/protobuf/proto"
    "net"
    "os"
    "time"
)
 
func main() {
    startTime := currentTimeMillis()
    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
 
//    transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "35998"))
	transport, err := thrift.NewTSocket(net.JoinHostPort("175.25.18.2", "35998"))
    if err != nil {
        fmt.Fprintln(os.Stderr, "error resolving address:", err)
        os.Exit(1)
    }
 
    useTransport := transportFactory.GetTransport(transport)
    client := rpc.NewProtoRpcServiceClientFactory(useTransport, protocolFactory)
    if err := transport.Open(); err != nil {
        fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
        os.Exit(1)
    }
    defer transport.Close()
 
	request := rpc.NewProtoRequest()
	request.TypeA1 = 101912
	var content UserRmbRpcDef.GetUserRmbRQ
	var userID uint64 = 9008000000005740
	content.UserId = &userID

	data, _ := proto.Marshal(&content)
	request.Contenet = string(data)
	fmt.Println(data)
	// test
	var con UserRmbRpcDef.GetUserRmbRQ
	err = proto.Unmarshal([]byte(request.Contenet), &con);
	if err == nil {
		fmt.Println("UserID:", *con.UserId)
	}
	
	fmt.Println("Content:", len(request.Contenet))
	
	r1, e1 := client.DealTwowayMessage(request)
	fmt.Println("Call->", r1, e1)
 
    endTime := currentTimeMillis()
    fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))
}
 
// 转换成毫秒
func currentTimeMillis() int64 {
    return time.Now().UnixNano() / 1000000
}