package main

import (
	"fmt"
	"time"
)

func timeout(t time.Duration) func() {
	start := time.Now()
	return func() {
		if time.Now().Sub(start) > t {
			panic("timeout")
		}
	}
}

func main() {
	defer timeout(time.Second)()

	time.Sleep(time.Second * 2)
	fmt.Println("hello, World")
}
