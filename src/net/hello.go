package main

import (
	"net"
	"os"
	"fmt"
	"bufio"
)

func panicTest() {
	panic(10)
	defer func() {
		v := recover()
		fmt.Println(v)
	}()
	//v := recover()
	//fmt.Println(v)
}

func main() {
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

	panicTest()
}
