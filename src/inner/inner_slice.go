package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []string{"w", "a", "g"}
	d := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("+%v", *d)

	fmt.Printf("%p", d)
}
