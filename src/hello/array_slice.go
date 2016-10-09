package main
import "fmt"
// 创建数组切片
// 基于数组
func create_slice_array() {
    var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var mySlice []int = myArray[5:]

    fmt.Println("Array: ")
    for _, v := range myArray { // ignore index
        fmt.Print(v, " ")
    }
    fmt.Println("\nSlice: ")
    for i, v := range mySlice {
        fmt.Print(v, "+", i, " ")
    }
    fmt.Println()

}

// 直接创建
func create_slice() {
    mySlice := make([]int, 5) 
    fmt.Println(mySlice)
    for i:=0; i<5; i++ {
        mySlice[i] = i
    }
    fmt.Println(mySlice)

    var mySlice2 []int = make([]int, 5, 10)
    fmt.Println(mySlice2)

    mySlice3 := []int{1, 2, 3, 4, 5}
    fmt.Println(mySlice3)

    for i, v := range(mySlice3) {
        fmt.Println("mySlice[", i, "] = ", v)
    }

    fmt.Println("len(mySlice2):", len(mySlice2))
    fmt.Println("cap(mySlice2):", cap(mySlice2))
    fmt.Println("cap(mySlice):", cap(mySlice))

    // 添加元素
    mySlice = append(mySlice2, 1, 2, 3)
    mySlice = append(mySlice2, mySlice3...)  // ...
    fmt.Println(mySlice)
    fmt.Println(mySlice2)
    fmt.Println("cap(mySlice):", cap(mySlice))
}

func array_slice() {
    array := [3]int {1, 2, 3}
    slice := []int {1, 2, 3}

    slice = make([]int, 5, 10)
    newSlice := slice[:6]
    fmt.Println(newSlice)

    fmt.Printf("%T %T\n", array, slice)
}

// 按其中较小的那个数组切片的元素个数进行复制
func slice_copy() {
    slice1 := []int {1, 2, 3, 4, 5}
    slice2 := make([]int, 5, 10)
    slice3 := make([]int, 3)
    copy(slice2, slice1)
    copy(slice3, slice1)

    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(slice3)
}

func main() {
    create_slice_array()

    create_slice()

    array_slice()

    slice_copy()
}
