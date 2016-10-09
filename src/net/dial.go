package main

import (
    "fmt"
    "net"
)
/*  Dial() 函数支持如下几种网络协议：
"tcp" 、 "tcp4" （仅限IPv4）、 "tcp6" （仅限IPv6）、
"udp" 、 "udp4" （仅限IPv4）、 "udp6" （仅限IPv6）、
"ip" 、 "ip4" （仅限IPv4）和 "ip6"（仅限IPv6）。
*/
/*
TCP链接：
conn, err := net.Dial("tcp", "192.168.0.10:2100")
UDP链接：
conn, err := net.Dial("udp", "192.168.0.12:975")
ICMP链接（使用协议名称）
conn, err := net.Dial("ip4:icmp", "www.baidu.com")
ICMP链接（使用协议编号）：
conn, err := net.Dial("ip4:1", "10.0.0.3")
*/
func main() {
    conn, err := net.Dial("tcp", "idu.com:80")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(conn)
    }
}
