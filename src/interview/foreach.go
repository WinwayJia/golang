package main

import (
	"fmt"
)

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// error
	for _, stu := range stus {
		//fmt.Printf("%p\n", &stu)
		m[stu.Name] = &stu
	}

	// right
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

}
