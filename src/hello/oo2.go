package main

import (
	"fmt"
	"strings"
)

type Part struct {
	Id   int
	Name string
}

func (p *Part) UpperCase() {
	fmt.Println("Uppercase")
	p.Name = strings.ToUpper(p.Name)
}

func (p Part) String() string {
	return fmt.Sprintf("id: %d name: %s", p.Id, p.Name)
}

type Item struct {
	id       uint64
	price    float64
	quantity int32
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item          // 嵌入
	catalogId int // 聚合
}

func main() {
	var p = &Part{1000, "json"}
	p.UpperCase()
	fmt.Println(p)

	var si = SpecialItem{Item{10000, 1.99, 20}, 20000}
	fmt.Println(si.Cost())

	si.price = 3.99
	fmt.Println(si.Cost())
}
