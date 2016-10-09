package  bubblesort

import "testing"
import "fmt"

func TestBubbleSort1(t *testing.T) {
    values := []int {5, 4, 3, 2, 1}
    BubbleSort(values)
    fmt.Println(values)
    if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
        values[4] != 5 {
        t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5")
    }
}

func TestBubbleSort2(t *testing.T) {
    src := [6]int {5, 5, 4, 2, 1, 3}
    values := src[0:]
    BubbleSort(values)

    for i:=0; i<len(values)-1; i++ {
        if values[i] > values[i+1] {
            t.Error("BubbleSort() failed. Got", src, "sorted: ", values)
        }
    }
}
