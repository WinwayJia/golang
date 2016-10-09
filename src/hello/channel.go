package main

import "fmt"
import "math/rand"
import "time"

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- rand.Intn(1024)
}

func TestMap(ch chan map[string] bool) {
	fmt.Println("TestMap")
	m := <- ch
	fmt.Println("", m["hello"])
}

func Funny() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <- ch
		fmt.Println("value recevied:", i)
	}

}

// 超时实现
func TestTimeout(ch chan int) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3e9)
		timeout <- true
	}()
	select {
	case <-ch:

	case <-timeout:
		fmt.Println("timeout...")
	}
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int, 0)
		go Count(chs[i])
	}
	for _, ch := range chs {
		val := <- ch
		fmt.Println("value =", val)
	}

	m := make(chan map[string] bool)
	mm := make(map[string] bool)
	mm["hello"] = true
	go TestMap(m)
	m <- mm

	// Funny()
	TestTimeout(chs[0])
	
//	var ch1 chan int // 双向
//	var ch2 <-chan int // 只读
//	var ch3 chan<- int // 只写
	
//	ch4 := make(chan int)
//	ch5 := <- chan int(ch4)  // 只读
//	ch6 := chan <- int (ch4) // 只写
	
}
