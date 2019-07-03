package main

import (
	"fmt"
)

func f() *int {
	x := 10
	y := 20

	fmt.Printf("x: %v\ny: %v\n", &x, &y)
	return &y
}

func main() {
	y := f()
	fmt.Printf("[main]y: %v\n", y)
}
