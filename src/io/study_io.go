package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	reader := strings.NewReader(`{"a": 1, "b": "hello"`)

	d := json.NewDecoder(reader)

	fmt.Println(d)
	var m map[string]interface{}
	d.Decode(&m)
	fmt.Println(m)

	buf, _ := ioutil.ReadAll(d.Buffered())
	fmt.Println(string(buf))
}
