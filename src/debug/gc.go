package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// Deal的定义
type DealTiny struct {
	Dealid    int32
	Classid   int32
	Mttypeid  int32
	Bizacctid int32
	Isonline  bool
	Geocnt    int32
}

func consume_memory() []int {
    buff := make([]int, 8192)
    return buff
}

func main() {
	const SIZE = 500000 // 50万
	m := make(map[int32]DealTiny, SIZE)
	var i int32
	for i = 0; i < SIZE; i++ { // 把数据放进内存
		m[i] = DealTiny{}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    // fmt.Println("request")
		// 模拟内存分配，做一些计算
		buffer := consume_memory()
		n := rand.Intn(4096) + 1024
		buffer = make([]int, n)
		for i := 0; i < n; i++ {
			buffer[i] = rand.Intn(1024)
		}
		c := 0
		for i := 0; i < n; i++ {
			if buffer[i] > 512 {
				c += 1
			}
		}
		fmt.Fprintf(w, "n: %d, more than 512 count: %d", n, c)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
