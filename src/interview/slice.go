package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
&{842350502824 2 3}
&{842350543376 5 6}
[1 2 10 100]
[2 10 100 200 300]
*/

func array() {
	array := [4]int{1, 2, 3, 4}
	slice := array[1:3]
	slice[1] = 10

	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Println(header)
	slice = append(slice, 100)
	slice = append(slice, 200)
	slice = append(slice, 300)
	fmt.Println(header)

	fmt.Println(array)
	fmt.Println(slice)
}

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	fmt.Println("\n\n")
	array()
}
