package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"
)

type Args struct {
	X, Y int
}

type Quotient struct {
	Quo, Rem int
}
type Arith int

func (i *Arith) Divide(arg *Args, quo *Quotient) error {
	if arg.Y == 0 {
		return errors.New("Divide by zero")
	}
	quo.Quo = arg.X / arg.Y
	quo.Rem = arg.X % arg.Y
	return nil
}

type StringRPC string

func (sr *StringRPC) Notify(msg string, ret *string) error {
	*ret = strings.ToUpper(msg)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	sr := new(StringRPC)
	err := rpc.Register(sr)
	if err != nil {
		log.Fatal("Register failed.", err)
		os.Exit(0)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error", e)
	}
	http.Serve(l, nil)
}
