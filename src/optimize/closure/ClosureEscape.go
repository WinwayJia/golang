package main

import (
	"fmt"
	"time"
)

func test() {
	x := 100
	go func(a int) {
		a++
		fmt.Println(a)
	}(x)
	x++
}

func ClosureTest() {
	y := 200
	go func() {
		y++
		fmt.Println(y)
	}()
	y++
}

func main() {
	test()
	ClosureTest()

	time.Sleep(time.Second)
}
