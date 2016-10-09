package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (ms *MyStruct)GetName() string {
	return ms.name;
}

func main() {
	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))

	var x float64 = 3.1415;
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))

	a := new(MyStruct)
	a.name = "godlikejia"
	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b[0])

	f := reflect.ValueOf(a).MethodByName("GetName").Call
	fmt.Printf("%v %v\n", a.GetName, f)
}


