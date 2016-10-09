package main

import (
	"testing"
)

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call(&Data{x: 100})
	}
}

func BenchmarkIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifaceCall(&Data{x: 100})
	}
}
