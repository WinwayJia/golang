package main

import (
	"fmt"
)

func main() {
	defer_call()

	defer_call2()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	//	panic("触发异常")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)

	return ret
}

/*
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
*/
func defer_call2() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
