package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d %v %v\n", *testNew(), makeSlice(), makeMap())
}

func testNew() *int {
	return new(int)
}

func makeSlice() *[]int {
	return new([]int)
}

func makeMap() *map[string]int {
	return new(map[string]int)
}
