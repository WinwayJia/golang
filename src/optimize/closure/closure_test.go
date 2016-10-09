package main

import (
	"testing"
)

func test(x int) int {
	return x * 2
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = test(i)
	}
}

func BenchmarkAnonymous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func(x int) int {
			return x * 2
		}(i)
	}
}

func BenchmarkClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = func() int {
			return i * 2
		}()
	}
}
