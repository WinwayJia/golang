package main

import "fmt"

type File struct {
	//
}

func (f *File) Read(buf []byte) (n int, err error) {
	fmt.Println("Read")
	return
}
func (f *File) Write(buf []byte) (n int, err error) {
	fmt.Println("Write")
	return
}

func (f *File) Seek(of int64, whence int) (pos int64, err error) {
	fmt.Println("Seek")
	return
}

func (f *File) Close() (err error) {
	fmt.Println("Close")
	return
}

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ISeeker interface {
	Seek(off int64, whence int) (pos int64, err error)
}

type ICloser interface {
	Close() error
}

// 将对象赋值给接口，要求该对象实现了接口的所有方法
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	// Add(b Integer)
}

/***********************************************/
// 将接口赋值给另一个接口
// 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值。
type ReaderWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

type IStream interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

func main() {
	fmt.Println("")
	var file1 IFile = new(File)
	buf := make([]byte, 1024)
	file1.Read(buf)
	// var file2 IReader = new(File)
	// var file3 IWriter = new(File)
	// var file4 ICloser = new(File)

	var a Integer = 1
	var b LessAdder = a
	// b.Add(10)
	fmt.Println(b.Less(0))

	// 接口查询
	// var file2 ReaderWriter = new(File)
	// if file3, ok := file2.(IStream); ok {
	//     fmt.Println("interface query")
	//     file3.
	// }

	// 类型查询

	// 接口组合

	// Any类型
	var v1 interface{} = 1
	var v2 interface{} = 10 //"abc"
	switch v2.(type) {
	case string:
		fmt.Println(v2, "is string")
	case int:
		fmt.Println(v2, "is int")
	}
	fmt.Println(v1, v2)
}
