package main

import (
	"fmt"
	"unsafe"
)

func slice2String() {
	slice := []byte{'1', '2', '3'}
	fmt.Println(string(slice))

	fmt.Println(*(*string)(unsafe.Pointer(&slice)))
}

func main() {
	slice2String()
}
