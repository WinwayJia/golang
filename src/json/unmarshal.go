package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var t struct {
		X int
	}

	err := json.Unmarshal([]byte(``), &t)
	fmt.Println(t.X, err)
}

func TestVimMode() {
	fmt.Println("Hello")
}
