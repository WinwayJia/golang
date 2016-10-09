package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

type SampleStruct struct {
	name string
	age  int
}

func (ss SampleStruct) String() string {
	return ss.name + " " + strconv.Itoa(ss.age)
}

type st struct {
	a byte
	b [8]byte
	c int
}

type Iface interface {
}

func main() {
	u := uint32(32)
	i := int32(1)
	fmt.Println(&u, &i) // 打印出地址
	p := &i             // p 的类型是 *int32
	//	p = &u              // &u的类型是 *uint32，于 p 的类型不同，不能赋值
	//	p = (*int32)(&u) // 这种类型转换语法也是无效的
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(*p)

	ss := SampleStruct{"Jack", 2}
	ss.age = 30
	fmt.Println(ss)

	t := st{a: 1, c: 3}
	fmt.Println(t)
	fmt.Println(unsafe.Pointer(&t.a))
	fmt.Println(unsafe.Pointer(&t.b))
	fmt.Println(unsafe.Pointer(&t.c))
	fmt.Println(unsafe.Sizeof(t.c))

	x := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("cap", cap(x[1:2]))
	fmt.Println("cap", cap(x[1:2:6]))

	var iface Iface = 10
	fmt.Println(&iface)
	if v, ok := iface.(int); ok {
		fmt.Println(iface.(int))
		fmt.Println(v)
	}

	//	fmt.Printf("%d\n", runtime.MemStats.Alloc/1024)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)

	func() {
		fmt.Println("closure")
	}()

	runtime.SetFinalizer(p, func(p *int32) {
		fmt.Println("GCing...")
	})

	const N = 100000
	big_obj := make([]*int, N)
	for i := 0; i < N; i++ {
		big_obj[i] = new(int)
	}
	//	runtime.GC()
	for i := 0; i < 50; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
				//				time.Sleep(time.Second)
			}
		}(i)
	}
	fmt.Println("GC begin...")
	runtime.GC()
	fmt.Println("GC end...")
	time.Sleep(10 * time.Microsecond)
}
