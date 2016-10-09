package main

import (
	"fmt"
	"net"
)

func main() {
	var ipStr = "127.0.0.1"
	addr := net.ParseIP(ipStr) // type IP []byte
	mask := addr.DefaultMask() // type IPMask []byte
	fmt.Printf("ip: %s, mask: %s\n", addr.String(), mask.String())

	/*
		type IPAddr struct {
		IP   IP
		Zone string // IPv6 scoped addressing zone
	}	*/
	var ipAddr *net.IPAddr

	ipAddr, err := net.ResolveIPAddr("ip", "baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", ipAddr.String())

	addrs, err := net.LookupHost("baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range addrs {
		fmt.Println("lookup:", v)
	}

	fmt.Println(net.LookupCNAME("map.baidu.com"))

	fmt.Println(net.LookupPort("tcp", "telnet"))
	fmt.Println(net.LookupPort("tcp", "http"))

	/*
		type TCPAddr struct {
			IP   IP
			Port int
			Zone string // IPv6 scoped addressing zone
		}*/
	tcpAddr, err := net.ResolveTCPAddr("tcp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tcpAddr)
}
