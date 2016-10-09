package main

import "fmt"
import "os"

func test(x int, y int) (a int, b int) {
    return x + 1, y - 1
}

func main() {
    var x = 10
    y := 100
    fmt.Printf("hello, world %d %d\n", x, y)

    fmt.Printf("%s\n", os.Args[1])
    
    x, y = test(10, 20);
    fmt.Printf("%d %d\n", x, y)

    var (
        a = 20
        b = 40
    )
    fmt.Printf("%d %d\n", a, b)

    // pointer
    var p* int = &a;
    *p = 1000
    fmt.Printf("%d %d %p\n", *p, a, p)

    const pi float64 = 3.14159265

    const (
        size uint64 = 10000
        eof = -1
    )

    fmt.Printf("%f %d %d\n", pi, size, eof);

    const (
        c1 = 1 << iota
        c2 = 1 << iota
        c3 = 1 << iota
    )
    fmt.Printf("%d %d %d\n", c1, c2, c3)

    var flag bool
    flag = true
    found := (1 == 2)

    fmt.Printf("result: %b %b\n", flag, found)

    // int32 and int not the same type
    var i1 int = 10
    i2 := 10
//    i2 = i1

    fmt.Printf("%d %d\n", i1, i2)
    if i1 == i2 {
        fmt.Printf("i1 equal i2\n")
    }

    var float_var1 float32;
    float_var1 = 12
    float_var2 := 12.0
    fmt.Printf("%f %.3f\n", float_var1, float_var2)

    var str string
    str = "Hello 世界"
    ch := str[0]
    fmt.Printf("the length of \"%s\" is %d, first char is %c\n", str, len(str), ch)

    // 以字节数组遍历
    for i:=0; i<len(str); i++ {
        fmt.Printf("%d --> %d\n", i, str[i])
        fmt.Println(i, "-->", str[i])
    }

    // 以unicode字符遍历
    for i, ch := range str {
        fmt.Println(i, ch)
    }

    test_array();

}

func test_array() bool {
    var byte_array [16]byte
    for i:=0; i<16; i++ {
        byte_array[i] = byte('a' + i)
    }
    fmt.Printf("%s\n", byte_array)

    for i, v := range byte_array {
        fmt.Println("byte_array[", i, "] = ", v)
    }

    return true;
}

func shuffle() bool {
    var cards [52]int
    for i:=0; i<len(cards); i++ {
        cards[i] = i
    }

    // shuffle
}
