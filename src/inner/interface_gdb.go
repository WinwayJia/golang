package main

import (
	"fmt"
	//	"io"
	"strconv"
)

type Stringer interface {
	String() string
	Get() uint64
}

func ToString(any interface{}) string {
	// z := any.(io.Reader) // runtime.assertE2I will panic
	//fmt.Println(z)
	if v, ok := any.(Stringer); ok { // runtime.assertE2I2
		var t Stringer = v
		fmt.Printf("type: %T %T\n", v, t)
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)

	}
	return "???"
}

type Binary uint64
type Binary32 uint32

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func (i Binary32) String() string {
	return strconv.FormatUint(uint64(i.Get()), 2)
}

func (i Binary32) Get() uint64 {
	return uint64(i)
}

func main() {
	b := Binary(200)
	b32 := Binary32(200)

	s := Stringer(b)
	s32 := Stringer(b32)

	any := (interface{})(b)

	fmt.Println(s.String())
	fmt.Println(s32.String())

	fmt.Println(ToString(any))

	var nothing interface{}

	nothing = uint64(200)
	fmt.Println(nothing)

	var face Stringer

	fmt.Println(face)

	var e interface{} = s
	fmt.Println(e)

	convert()
}

// Iface Eface Type
func convert() {
	fmt.Println("convert: ")
	t := Binary(100)
	// Type -> Iface
	i := Stringer(t)
	// Iface -> Type
	// t = (Binary)(i)// cannot convert i (type Stringer) to type Binary: need type assertion
	t = i.(Binary)
	fmt.Println(t, i)

	// Type -> Eface
	var e interface{} = t
	// Eface -> Type
	t = e.(Binary) // cannot convert e (type interface {}) to type Binary: need type assertion
	fmt.Println(t, e)

	// Iface -> Eface
	e = i
	// Eface -> Iface
	i = e.(Stringer) // cannot use e (type interface {}) as type Stringer in assignment:
	fmt.Println(i, e)
}
