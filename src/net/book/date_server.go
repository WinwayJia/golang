package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func handleClient(conn *net.TCPConn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}

func main() {
	addr := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept failed.", err)
			continue
		}
		go handleClient(conn)
	}
}
