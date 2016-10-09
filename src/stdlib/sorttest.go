package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [...]int{2, 4, 3, 1, 5}
	sort.Ints(arr)
	fmt.Println(arr)
}
