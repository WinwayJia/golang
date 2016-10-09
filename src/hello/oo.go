package main
import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
    return a < b
}

func (a Integer) Add1(b Integer) {
    a += b
}

func (a *Integer) Add2(b Integer) {
    *a += b
}

type Rect struct {
    x, y float64
    width, height float64
}

func (r *Rect) Area() float64 {
    return r.width * r.height
}

// 继承
type Base struct {
    Name string
}

func (b *Base) Foo() {
    fmt.Println("Base::Foo")
}

func (b *Base) Bar() {
    fmt.Println("Base::Bar")
}

func (b *Base) Test() {
    fmt.Println("Base::Test")
}

type Foo struct {
    Info Base
    Age  int
    *Base
}

func (f *Foo) Foo() {
    fmt.Println("Foo:Foo")
}

func (f* Foo) Bar() {
    fmt.Println("Foo:Bar")
}


func main() {
    var a Integer = 1
    if a.Less(2) {
        fmt.Println(a, "less 2")
    }
    arr := [...]int {1, 2, 3}
    arr2 := arr
    arr[1] = 4

    fmt.Println(arr)
    fmt.Println(arr2)

    a.Add1(10)
    fmt.Println(a)
    a.Add2(10)
    fmt.Println(a)

    rect := &Rect{x:0, y:0, width: 10, height: 10}
    fmt.Println("Area:", rect.Area())
    rect1 := new(Rect)
    fmt.Println("Area:", rect1.Area())

    var b Base;
    b.Foo()
    b.Bar()

    var f Foo;
    f.Foo()
    f.Bar()
    f.Info.Bar()
    f.Test()
}
