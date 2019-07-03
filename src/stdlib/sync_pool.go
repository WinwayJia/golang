package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	p.Put(1)
	b := p.Get().(int)
	c := p.Get().(int)
	fmt.Println(a, b, c)

	a = p.Get().(int)
	p.Put(1)
	runtime.GC() //手动调用GC
	b = p.Get().(int)
	fmt.Println(a, b)
}
