package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// 通过unsafe.Pointer来转化类型
func typeConv() {
	var a int64 = 3
	var b float64 = float64(a)
	fmt.Println(&a) // 0xc42000e248
	fmt.Println(&b) // 0xc42000e250

	//a = float64(a) // cannot use float64(a) (type float64) as type int64 in assignment
	b = *(*float64)(unsafe.Pointer(&a))

	fmt.Println(a, b)

	var n int64 = 3
	var pn = &n                             // n的指针
	var pf = (*float64)(unsafe.Pointer(pn)) // 通过Pointer来将n的类型转为float
	fmt.Println(*pf)                        // 2.5e-323
	*pf = 3.5
	fmt.Println(n) // 4615063718147915776

	fmt.Println(pf) // 0xc42007a050
	fmt.Println(pn) // 0xc42007a050
}

// 通过uintptr来计算偏移量

func offset() {
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])                               // index为1的元素
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0])) // 拿到index为3的指针
	*(*int)(p3) = 4                                           // 重新赋值
	fmt.Println(a)                                            // a = [0 1 2 4]

	s := strconv.Itoa(10000)
	ps := unsafe.Pointer(&s)
	length := *(*int)(unsafe.Pointer((uintptr(ps) + 8)))
	fmt.Println("length: ", length)
	//	pb := unsafe.Pointer(uintptr(bytes) + 5)
	//	*(*byte)(pb) = 'n'
	//pb := unsafe.Pointer(uintptr(bytes) + 8)
	//*(*int)(pb) = 6

	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	header.Len = 6
	ptr := unsafe.Pointer(header.Data)

	*(*byte)(ptr) = '2'

	fmt.Printf("%v %s %d %d %c\n", header, s, len(s), ptr, *(*byte)(ptr))
}

func string2bytes(s string) []byte {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh := reflect.SliceHeader{
		Data: header.Data,
		Len:  header.Len,
		Cap:  header.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sh))
}

func bytes2string(b []byte) string {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: header.Data,
		Len:  header.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}

func main() {
	var x uintptr
	fmt.Println(unsafe.Sizeof(x)) // 8
	offset()

	slice := string2bytes("hello")
	fmt.Printf("slice: %T %v\n", slice, slice)

	s := bytes2string(slice)
	fmt.Printf("slice: %T %v\n", s, s)
}
