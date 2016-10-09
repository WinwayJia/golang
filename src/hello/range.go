package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := map[int]string{1: "1", 2: "2"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	a := [3]int{1, 2, 3}
	for i, v := range a { // range会复制对象，index、value 都是从复制品中取出
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(unsafe.Pointer(&a[1]), a)
		}
		//		a[i] = v + 100  // v = 1, 2, 3
		a[i] += 100
	}

	fmt.Println(unsafe.Pointer(&a[1]), a)
}
