package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())

	ch := make(chan string, 10)

	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				ch <- "new msg"
				fmt.Println("send msg")
			}
		}
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("recvMsg: ", msg)
		case <-time.After(5 * time.Second):
			fmt.Println("timeout:", time.Now().Unix())
		}
	}
}
