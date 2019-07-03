package main

import (
	"fmt"
)

// 注意这里的 //go:noinline 编译器指令。。不要省略掉这部分
//go:noinline
func add(a, b int32) (int32, bool) {
	var c int32
	c = a + b
	a = c * c
	return a + b, true
}

func main() {
	add(10, 32)

	caller()
}

// Go 的调用规约要求每一个参数都通过栈来传递，这部分空间由 caller 在其栈帧(stack frame)上提供。
func caller() {
	var w int64
	var x int32 = 1
	y := 2
	z := 3
	fmt.Printf("%p %p %p %p\n", &w, &x, &y, &z)
	callee(x, y)
	a := 1
	b := 2
	fmt.Printf("%p %p %p %p %p %p\n", &w, &x, &y, &z, &a, &b)
}

func callee(x int32, y int) (ret int) {
	fmt.Printf("%p %p %p\n", &x, &y, &ret)

	return
}
