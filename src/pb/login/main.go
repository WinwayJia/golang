package main

import (
	"log"
	"net"
)

func dealClient(conn *net.Conn) {
	client := &XClient{client:conn}

	client.Login()
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {

	}
	log.Println("")
	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println("Accept failed.", err.Error())
			continue
		}
		dealClient(&client)
	}

}
