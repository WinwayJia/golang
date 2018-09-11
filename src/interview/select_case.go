package main

// 使用SelectCase来做不定数量的channel选择
import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var n int = 1
	var chs = make([]chan int, n)

	var worker = func(n int, c chan int) {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}

	//不定数量的channel数组
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go worker(3+i, chs[i])
	}

	var selectCase = make([]reflect.SelectCase, n)
	//将channel绑定到SelectCase
	for i := 0; i < n; i++ {
		selectCase[i].Dir = reflect.SelectRecv //设置信道是接收,可以为下面值之一
		/*
			const (
				SelectSend    SelectDir // case Chan <- Send
				SelectRecv              // case <-Chan:
				SelectDefault           // default
			)
		*/
		selectCase[i].Chan = reflect.ValueOf(chs[i])
	}

	numDone := 0
	//从所有channel中取出最先到达的N个值
	for numDone < n {
		chosen, recv, recvOk := reflect.Select(selectCase)
		if recvOk {
			fmt.Println(chosen, recv.Int(), recvOk)
			numDone++
		} else {
			fmt.Println("recv error")
		}
	}

	MyUsage()
}

func MyUsage() {
	fmt.Println("MyUsage")

	ch := make(chan int, 1)
	selectCase := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(time.After(5 * time.Second)),
		},
	}
	ch <- 1

	chosen, recv, ok := reflect.Select(selectCase)
	fmt.Println(chosen, recv, ok)
}
