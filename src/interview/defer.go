package main

/*
The behavior of defer statements is straightforward and predictable. There are three simple rules:

1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	return
}

2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i) // 3210
	}
}

3. Deferred functions may read and assign to the returning function's named return values.
func c() (i int) {
	defer func() { i++  }()
	return 1   // this function returns 2
}
*/

import (
	"fmt"
)

func main() {
	defer_call()

	defer_call2()

	fmt.Println("DeferFunc*(1)")
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
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

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
