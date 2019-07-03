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

func nilSlice() {
	// nil slice
	var nil1 []int
	var nil2 = *new([]int)
	fmt.Println("nil slice: ", nil1 == nil, nil2 == nil)

	// empty slice
	e1 := make([]int, 0, 0)
	var e2 = []int{}
	fmt.Println("empty slice: ", e1 != nil, e2 != nil)
}

func sliceInit() {
	s1 := []int{0, 1, 2, 3, 8: 100} // 索引号
	fmt.Println(s1, len(s1), cap(s1))
}

func expandSlice(s []int) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(header)
	s = append(s, []int{1, 2, 3, 4, 5, 6, 7, 8}...)
	header = (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(header)
	fmt.Println(s)
}

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	fmt.Println("\n\n")
	array()

	nilSlice()

	sliceInit()

	s1 := make([]int, 0, 1)
	expandSlice(s1)
	fmt.Println(s1)
}
