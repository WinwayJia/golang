package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

func main() {
	var s Shaper
	fmt.Println("value of s is ", s)
	fmt.Printf("type of s is %T\n", s)
}
