package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p *People) String() string {
	fmt.Println(1)
	return fmt.Sprintf("print: %v", p) // p will String, 无限递归
}

type Student struct {
	Name string
}

func main1() {
	m := map[string]Student{"people": {Name: "zhoujielun"}}
	fmt.Println(m)
	//m["people"].Name = "wuyanzu"
	x := &m["people"]
	x.Name = "wuyanzu"
	fmt.Println(m)
}

func main() {
	//p := &People{}
	//	p.String()
	main1()
}
