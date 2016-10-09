package ipc
import (
    "testing"
)

type EchoServer struct {

}

func (server* EchoServer) Handle(method, params string) *Response {
    var resp Response

    resp.Code = "OK"
    resp.Body = "ECHO:From " +  params

    return &resp
}

func (server* EchoServer) Name() string {
    return "EchoServer"
}

func TestIpc(t *testing.T) {
    server := NewIpcServer(&EchoServer{})
    client1 := NewIpcClient(server)
    client2 := NewIpcClient(server)

    resp1, err := client1.Call("ECHO", "Client1")
    if err != nil {
        return
    }
    resp2, err := client2.Call("ECHO", "Client2")
    if err != nil {
        return
    }

    if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
        t.Error("IpcClient.Call failed. resp1:", resp1.Body, "resp2:", resp2.Body)
    }

    client1.Close()
    client2.Close()
}
