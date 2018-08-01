package main

import (
	"fmt"
	"sync"
	"time"
)

type threadSafeSet struct {
	sync.RWMutex

	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {

	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()
	}()

	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}

	v := <-th.Iter()

	fmt.Sprintf("%s%v", "ch", v)

	time.Sleep(3 * time.Second)

	main1()
}

func main1() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
