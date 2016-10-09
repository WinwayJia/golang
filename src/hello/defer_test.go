// go test -bench="." defer_test.go
package lib

import (
	"sync"
	"testing"
)

var lock sync.Mutex

func test() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
//	lock.Unlock()
	defer lock.Unlock()
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}

/*
BenchmarkTest           50000000                32.3 ns/op
BenchmarkTestDefer      10000000               168 ns/op
*/
