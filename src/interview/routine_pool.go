package main

import (
	"fmt"
	"time"
)

func maxGoRoutine() {
	num := 0
	for {
		go func() {
			for {
				time.Sleep(1)
			}
		}()
		num++
		if num%10000 == 0 {
			fmt.Println(num / 10000)
		}
	}
}

func main() {
	maxGoRoutine()
}
