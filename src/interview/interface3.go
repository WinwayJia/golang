package main

import "fmt"

type Iname interface {
	Mname()
}
type St struct{}

func (St) Mname()         {}
func (St) String() string { return "string" }

func foo(i interface{}) {
	s, ok := i.(Iname)
	if ok {
		fmt.Printf("%T ==> %T success\n", i, s)
		s.Mname()
	}

	st, ok := i.(St)
	if ok {
		fmt.Printf("%T ==> %T success\n", i, st)
	}
}

func main() {
	var t *St
	if t == nil {
		fmt.Println("t is nil") // print
	} else {
		fmt.Println("t is not nil")
	}

	var i Iname = t
	fmt.Printf("%T\n", i) // *main.St

	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is not nil") // print
	}

	fmt.Printf("i is nil pointer:%v\n", i == (*St)(nil)) // true

	var st St
	i = st
	i.Mname()

	foo(i)
}
