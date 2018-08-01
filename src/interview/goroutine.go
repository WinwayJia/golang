package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("first i: ", i)
			wg.Done()
		}()
		for j := 0; j < 10000; j++ {
			fmt.Println("first sleep")
		}
		//time.Sleep(time.Microsecond)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("second i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Microsecond)
}
