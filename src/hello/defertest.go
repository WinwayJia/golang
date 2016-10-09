package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer close 1")
	defer fmt.Println("defer close 2")
	defer func() {
		defer fmt.Println("defer in func")
	}()
	fmt.Println("End of main")
}
