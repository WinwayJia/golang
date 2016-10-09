package main

import "fmt"
import "math/rand"
import "runtime"
import "sync"
// import "time"
// Go语言包中的 sync 包提供了两种锁类型： sync.Mutex 和 sync.RWMutex
var lock sync.Mutex

func Add(ch chan int, index int) {
	var x int
	var y int
	for i := 0; i < 20; i++ {
		lock.Lock()
		x = rand.Intn(65536)
		y = rand.Intn(65536)
		fmt.Println(index, "x + y =", x+y)
		lock.Unlock()
		// time.Sleep(1e9)
	}
	ch <- index
}

// 全局唯一性操作
var once sync.Once
func setup() {  // no argument
	fmt.Println("setup func", x)
}

func doprint() {
	once.Do(setup)
	fmt.Println("doprint")
}

func main() {
	fmt.Println("go routine")
	chs := make([]chan int, 2)
	for i := 0; i < 2; i++ {
		chs[i] = make(chan int, 1)
		go Add(chs[i], i)
	}

	for i := 0; i < 2; i++ {
		<- chs[i]
	}
	fmt.Println(runtime.NumCPU())

	doprint()
	doprint()
}
