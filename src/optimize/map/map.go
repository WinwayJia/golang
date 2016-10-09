package main

import (
	"runtime"
)

const capacity = 500000

var d interface{}

func value() interface{} {
	m := make(map[int]int, capacity)
	for i := 0; i < capacity; i++ {
		m[i] = i
	}
	return m
}

func pointer() interface{} {
	m := make(map[int]*int, capacity)
	for i := 0; i < capacity; i++ {
		v := i
		m[i] = &v
	}
	return m
}

func main() {
	//	_ = value()
	_ = pointer()
	for i := 0; i < 20; i++ {
		runtime.GC()
	}
}
