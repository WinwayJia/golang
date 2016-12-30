package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"unsafe"
	"reflect"
)

var (
	TestError = fmt.Errorf("test error type. %d", 10)
	Test2Error = errors.New("test error type")
)

func writeFile() error {
	file, err := os.OpenFile("readme.txt", os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	v := 0
	val := 65 / v
	ptr := unsafe.Pointer(&val)
	str := (*byte)(ptr)
	fmt.Println(*str, reflect.TypeOf(*str))
	//file.Write([]byte(ptr))
	return nil
}

func readFile() {

}

func main() {
	reader := strings.NewReader(`{"a": 1, "b": "hello"`)

	d := json.NewDecoder(reader)

	fmt.Println(d)
	var m map[string]interface{}
	d.Decode(&m)
	fmt.Println(m)

	buf, _ := ioutil.ReadAll(d.Buffered())
	fmt.Println(string(buf))

	writeFile()
}
