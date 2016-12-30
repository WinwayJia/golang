package main

import (
	"net"
)

type XClient struct {
	UserID int64
	client *net.Conn
}

func (c *XClient) Login() {

}

func (c *XClient) Ping() {

}
