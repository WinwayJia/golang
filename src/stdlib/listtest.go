package main

import (
	//	"container/list"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(read io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := read.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func ReadFromString() {
	reader := strings.NewReader("Hello")
	data, err := ReadFrom(reader, 20)
	if err != nil {
		fmt.Println("read failed.", err.Error())
	} else {
		fmt.Println(string(data))
	}
}

func ReadFromStdin() {
	data, err := ReadFrom(os.Stdin, 20)
	if err != nil {
		fmt.Println("read failed.", err.Error())
	} else {
		fmt.Println(string(data))
	}
}

func ReadFromFile() {
	file, err := os.Open("listtest.go")
	if err != nil {
		return
	}
	data, err := ReadFrom(file, 100)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}

type Test struct {
	Name string
}

//func (t *Test) String() string {
//	return "Test::string"
//}

func (t Test) String() string {
	return t.Name
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	t := Test{"Hello, Test"}
	tt := &Test{"Hello, Test"}
	fmt.Println(&t, tt)

	//	var ll list.List
	//	ll.Init()
	//	ll := list.New()
	//	ll.PushBack(123)
	//	ll.PushBack("abc")

	//	ll.PushFront("Front")
	//	front := ll.Front()
	//	for ; front != nil; front = front.Next() {
	//		fmt.Println(front.Value)
	//	}

	//	ReadFromString()
	//	ReadFromStdin()
	//	ReadFromFile()

	//	var a Integer = 1
	//	var b1 LessAdder = &a //OK
	//	var b2 LessAdder = a  //not OK
}
