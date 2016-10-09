package main

import (
	"errors"
	"fmt"
	"log"
	"net/rpc"
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

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	args := &Args{20, 7}
	quo := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quo, nil)
	reply := <-divCall.Done

	fmt.Printf("%d %d\n", quo.Quo, quo.Rem)
	fmt.Println(reply)

	ret := new(string)
	call := client.Go("StringRPC.Notify", "hello", ret, nil)
	<-call.Done
	fmt.Println(*ret)
}
