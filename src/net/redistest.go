package main

import (
    "fmt"
    "net"
    "os"
    "bytes"
    "time"
)

type RedisClient interface {
    Connect() error
    Get(key string) (ret string, err error)
}

type SimpleRedisClient struct {
    Conn  net.Conn
    Addr  string
}

func (client *SimpleRedisClient) Connect() (err error) {
    client.Conn, err = net.Dial("tcp", client.Addr)
    if err != nil {
        return err
    }
    return nil
}

func (client *SimpleRedisClient) Get(key string) (ret string, err error)  {
    str := fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(key), key)

    _, err = client.Conn.Write([]byte(str))
    if err != nil {
        fmt.Println("dial error", err.Error)
        os.Exit(1)
    }

    var data [512]byte
    n, err := client.Conn.Read(data[0:])
    if err != nil {
        fmt.Println("read error", err)
        os.Exit(1)
    }

    result := bytes.NewBuffer(nil)
    result.Write(data[0:n])
    ret = string(result.Bytes())

    return ret, nil
}

type TestRedisClient struct {

}

func (client* TestRedisClient) Connect() error {
    return nil
}

func (client* TestRedisClient) Get(key string) (str string, err error) {
    fmt.Println("TestRedisClient::Get")
    return "Test", nil
}

func NewRedisClient() *SimpleRedisClient {
    return &SimpleRedisClient{Addr: "127.0.0.1:6379"}
}

func main() {
    var client RedisClient
    // client = &SimpleRedisClient{Addr: "127.0.0.1:6379"}
    client = NewRedisClient()
    // client = &TestRedisClient{}
    client.Connect()
    ret, err := client.Get("Test")
    if err != nil {
        os.Exit(1)
    }
    fmt.Printf("ret:%s\n", ret)

    time.Sleep(30 * time.Second)
}
