package main

import (
	"runtime"
	"time"
	"unsafe"
)

type data struct {
	x [1024 * 1024]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func main() {
	const N = 10000
	cache := new([N]uintptr)

	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}

	runtime.SetFinalizer()
}
