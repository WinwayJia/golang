package main

import "fmt"

func DeferFunc() {
	fmt.Println("hello world")
}
func NoDefer() {
	DeferFunc()
}

func Defer() {
	defer DeferFunc()
}

func main() {
	NoDefer()
	Defer()
}
