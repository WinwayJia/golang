package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string
}

type Admin struct {
	User
	title string
}

func main() {
	var x float64 = 10.32
	v := reflect.ValueOf(x)
	fmt.Printf("type: %v\n", reflect.TypeOf(x))
	fmt.Printf("type: %T\n", x)

	fmt.Println(v)

	var admin Admin
	t := reflect.ValueOf(admin)
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println("f", f.Type())
	}
}
