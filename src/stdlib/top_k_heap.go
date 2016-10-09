package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"math/rand"
	"strconv"
)

type Type []int

func (t *Type) Len() int {
	return len(*t)
}

func (t *Type) Less(i, j int) bool {
	return (*t)[i] > (*t)[j]
}

func (t *Type) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *Type) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *Type) Pop() (v interface{}) {
	*t, v = (*t)[:len(*t)-1], (*t)[len(*t)-1]
	return
}

func (t *Type) String() string {
	var buffer bytes.Buffer
	for _, v := range *t {
		buffer.WriteString(strconv.Itoa(v) + " ")
	}
	return buffer.String()
}

func main() {
	fmt.Println("stdlib heap")
	v := new(Type)
	//	for i := 0; i < 5; i++ {
	//		v.Push(i)
	//	}
	//	heap.Init(v)
	//	fmt.Println(v)

	//	//	v.Push(10)
	//	heap.Push(v, 10)
	//	heap.Push(v, 5)
	//	fmt.Println(v)
	//	_ = heap.Pop(v)
	//	fmt.Println(v)

	// top k
	k := 3
	for i := 0; i < 10; i++ {
		rand := rand.Intn(65536)
		fmt.Println(rand)
		if v.Len() < k {
			fmt.Println("not enough")
			heap.Push(v, rand)
		} else {
			heap.Push(v, rand)
			tmp := heap.Remove(v, k)
			fmt.Println("tmp", tmp)
		}
	}
	fmt.Println(v)
	for i := 0; i < k; i++ {
		fmt.Println(heap.Remove(v, 0))
	}
}
