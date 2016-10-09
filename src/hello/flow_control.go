package main

import "fmt"

// 在有返回值的函数中，不允许将“最终的” return 语句包含在 if...else... 结构中
func test_if() {
    var m int
    m = 10
    n := 20
    if (n > m) {
        fmt.Println("n > m")
    } else if (n < m) {
        fmt.Println("n < m")
    } else {
        fmt.Println("n == m")
    }
}

func test_switch(i int) {
    switch i {
        case 0:
            fmt.Println("0")
            fallthrough
        case 1:
            fmt.Println("1")
        case 4, 5, 6:
            fmt.Println("4 5 6")
        default:
            fmt.Println("default")
    }

    switch {
        case i>=0 && i <= 3:
            fmt.Println("[1, 3]")
        case i>3 && i <= 6:
            fmt.Println("(3, 6]")
        default:
            fmt.Println("default")
    }

    fmt.Println("\n\n")
}

// just for, no while, do while
func test_for() {
    slice := make([]int, 10)
    for i, _ := range slice {
        slice[i] = i 
    }
    sum := 0
    for _, v := range slice {
        sum += v
    }

    fmt.Printf("sum = %d\n", sum)

    // dead loop
    for {
        sum ++
        if sum > 100 {
            break
        }
    }
    fmt.Printf("sum = %d\n", sum)

    /*
    fmt.Println(slice)
    for m, n := 0, len(slice); m < n; m = m + 1, n = n - 1 {
        slice[m], slice[n] = slice[n], slice[m]
    }
    fmt.Println(slice)
    */

    /*
    count := 10
    for m := 0; m < count; m++ {
        for n := 0; n < count; n++ {
            if m * n > 40 {
                break MLoop
            }
        }
    }
    MLoop:
    fmt.Println("break outer loop")
    */
}

func test_goto() {
    i := 0
    HERE:
    fmt.Print(i, " ")
    i ++
    if i < 10 {
        goto HERE
    }
    fmt.Println()
}

func main() {
    test_if()
    for i:=0; i<10; i++ {
        test_switch(i)
    }

    test_for()

    test_goto()
}
