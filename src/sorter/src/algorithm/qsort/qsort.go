package qsort

func QuickSort(values []int) {
    quickSort(values, 0, len(values))
}

func quickSort(values []int, left int, right int) {
    if left >= right {
        return
    }
    prime := values[left]
    j := left + 1
    for i := left + 1; i < right; i++ {
        if values[i] < prime {
            values[i], values[j] = values[j], values[i]
            j ++
        }
    }

    values[j-1], values[left] = values[left], values[j-1]

    quickSort(values, left, j-1)
    quickSort(values, j, right)
}

