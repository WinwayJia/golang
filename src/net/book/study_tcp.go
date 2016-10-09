package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	var conn *net.TCPConn
	addr, err := net.ResolveTCPAddr("tcp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	//	buff := make([]byte, 2048)
	//	n, err := conn.Read(buff)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	buff, err := ioutil.ReadAll(conn)
	fmt.Println(string(buff))
}
