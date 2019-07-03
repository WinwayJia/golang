package main

// https://research.swtch.com/interfaces

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 32)
	}
	return "???"
}

type Binary uint64

type Getter interface {
	Get() uint64
}

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

type Setter interface {
	Set(x int)
}

type st struct {
	x int
}

func (s *st) Set(x int) {
	s.x = x
}

func (s st) String() string {
	return "st"
}

func main() {
	b := Binary(4294967295)
	s := Stringer(b)
	g := Getter(b)

	s.String()
	g.Get()

	fmt.Println("Hello go interface", s)

	copyCase()
}

func copyCase() {
	s := &st{x: 10}
	str := Stringer(s) // this case a copy
	setter := Setter(s)

	str.String()
	setter.Set(1000)
}
