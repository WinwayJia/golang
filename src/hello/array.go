package main
//import "fmt"
import (
    "fmt"
)
/*
###数组是值类型###
需要特别注意的是，在Go语言中数组是一个值类型（value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。
如果将数组作为函数的参数类型，则在函数调用时该参数将发生数据复制。因此，在函数体中无法修改传入的数组的内容，因为函数内操作的只是所
传入数组的一个副本。
*/

func modify(array [10]int) {
    array[0] = 10
    fmt.Println("in modify", array)
}

func main() {
    var array [10]int
    for i:=0; i<len(array); i++ {
        array[i] = i
    }
    modify(array)

    fmt.Println("in main", array)

    // 不能make array，slice可以
    arr := make([]int, 5)
    fmt.Println(arr)
}
