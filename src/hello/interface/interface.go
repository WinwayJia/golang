package main

import (
	"bytes"
	"fmt"
)

type S struct {
	i int32
}

func (p *S) Get() int32  { return p.i }
func (p *S) Put(i int32) { p.i = i }

type I interface {
	Get() int32
	Put(int32)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(1)
	if str, ok := p.(I); ok {
		fmt.Println("p is string..", str)
	}
}

func g(i interface{}) int32 { return i.(I).Get() }

func h() {
	var s S
	fmt.Println(g(&s))
	fmt.Println(g(s))
}

func concat(items ...interface{}) string {
	var buf bytes.Buffer
	for _, item := range items {
		switch v := item.(type) {
		case int, int32, uint32, int64, uint64:
			fmt.Println(v)
			buf.WriteString(fmt.Sprintf("%d ", v))
		case string:
			buf.WriteString(v)
		default:
			buf.WriteString("?")
		}
	}

	return buf.String()
}

func main() {
	var s S
	f(&s)
	//h()

	var x int = 10
	var y int32 = 20
	var z int64 = 30
	var str = "string"
	var b bool = true

	fmt.Println("concat", concat(x, y, z, str, b))
}
