package main

import (
	"fmt"

	"./mymath"
)

// 不支持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)。
//• 无需声明原型。
//• 支持不定长变参。
//• 支持多返回值。
//• 支持命名返回参数。
//• 支持匿名函数和闭包。

// 形如 ...type 格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数。 变参本质上就是 slice
func myfunc3(arg ...int) {
	sum := 0
	for _, v := range arg {
		sum += v
	}
	fmt.Println("sum =", sum)
}

func myfunc(args ...int) {
	// 按原样传递
	myfunc3(args...)
	// 传递片段，实际上任意的int slice都可以传进去
	myfunc3(args[1:]...)
}

func myPrintf(args ...interface{}) (sum int, str string) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			sum += arg.(int)
		case string:
			str += arg.(string)
		}
	}
	return
}

func test(fn func() int) int {
	return fn()
}

type FormatString func(string, int, int) string

func add(x, y int) (z int) {
	z = x + y
	return z
}
func main() {
	ret, err := mymath.Add(10, 20)
	fmt.Println(ret, " ", err)

	ret, err = mymath.Add(-10, 20)
	fmt.Println(ret, " ", err)

	myfunc(10, 20, 30)
	myfunc(10, 20, 30, 40, 50)

	num, str := myPrintf(10, "1", 2, "2", "Hello")
	fmt.Println(num)
	fmt.Println(str)

	// 匿名函数
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println(f(10, 20))

	//	匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。
	fn := func() { fmt.Println("匿名函数") }
	fn()
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(fns[0](10), fns[1](10))

	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello" },
	}
	fmt.Println(d.fn())

	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, chan func" }
	fmt.Println((<-fc)())

}
