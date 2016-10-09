package qsort
import "testing"

func TestQuickSort1(t *testing.T) {
    values := [5]int {2, 3, 4, 1, 5}
    QuickSort(values[0:])

    for i := 0; i < len(values) - 1; i++ {
        if values[i] > values[i+1] {
            t.Error("sorted1, result", values);
        }
    }
}

func TestQuickSort2(t *testing.T) {
    values := [6]int {2, 2, 3, 5, 1, 4}
    QuickSort(values[0:])

    for i := 0; i < len(values) - 1; i++ {
        if values[i] > values[i+1] {
            t.Error("sorted2, result", values);
        }
    }
}
