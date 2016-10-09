package main

import (
	//	"bufio"
	//	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	strReader := strings.NewReader("Hello")
	//	bufReader := bufio.NewReader(strReader)

	//	data, _ := bufReader.Peek(5)
	//	data = make([]byte, 10)
	//	n, _ := bufReader.Read(data)
	//	fmt.Println(data, n, bufReader.Buffered())

	//	io.Copy(os.Stdout, bufReader)
	io.Copy(os.Stdout, strReader)
}
