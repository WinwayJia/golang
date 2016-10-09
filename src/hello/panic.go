package main

import "fmt"

func test() {
	defer func() {
		/*
		if err := recover(); err != nil {
			fmt.Println(err.(string))	// 
		}
		*/
	}()

	panic("painc error")
}

func main() {
	fmt.Println("test panic and recover")
	test()
}
