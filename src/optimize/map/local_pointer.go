package main

import (
	"fmt"
	"unsafe"
)

func test() *int { // ok, x will escapes to heap
	x := 20
	return &x
}

func test_value() int {
	x := 20
	return x
}

func main() {
	p := test()
	v := test_value()
	fmt.Println(p, *p, v)

	d := struct {
		s string
		x int
	}{"abc", 10}

	// 指针运算
	p1 := uintptr(unsafe.Pointer(&d))
	p1 += unsafe.Offsetof(d.x)
	p2 := unsafe.Pointer(p1)
	px := (*int)(p2)
	*px = 20
	fmt.Println(d)

	// 指针数组
	a, b, c := 1, 2, 3
	var arrayOfPointer [3]*int = [3]*int{&a, &b, &c}
	fmt.Println(arrayOfPointer)
	// 指向数组的指针
	var pointerToArray *[3]int = &[3]int{1, 2, 3}
	fmt.Println(pointerToArray)
}
