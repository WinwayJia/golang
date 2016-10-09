package main

import (
	"os"
	"fmt"
	"flag"
	"net"
)

var dest_ip *string = flag.String("ip", "127.0.0.1", "dest ip addr")
var dest_port *uint = flag.Uint("port", 8080, "dest port")

func proxy_return(front, back net.Conn) {
	data := make([]byte, 512)
	for {
		n, err := back.Read(data)
		if n == 0 {
			break
		}
		fmt.Println("return", string(data[:n]), n, err)
		if err == nil {
			front.Write(data[:n])
		}
	}
	fmt.Println("exiting proxy_return routine")
}

func proxy_it(front net.Conn) {
	addr := fmt.Sprintf("%s:%d", *dest_ip, *dest_port)
	fmt.Println("Address:", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dial error. %s", err.Error())
		os.Exit(1)
	}
	
	defer conn.Close()
	defer front.Close()
	
	go proxy_return(front, conn)
	
	data := make([]byte, 512)
	for {
		n, err := front.Read(data)
		if n == 0 {
			break
		}
		fmt.Println(string(data[:n]))
		if err == nil {
			conn.Write(data[:n])
		}
	}
	fmt.Println("exiting proxy_it routine")
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen failed.", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err == nil {
			fmt.Printf(conn.RemoteAddr().String())
			go proxy_it(conn)
		}
	}
}