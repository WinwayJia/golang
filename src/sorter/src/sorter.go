package main

import "flag"
import "fmt"
import "bufio"
import "os"
import "io"
import "strconv"
import "time"
import "algorithm/bubblesort"
import "algorithm/qsort"

var infile = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")
var ip = flag.Int("flagname", 1234, "help message for flagname")

func readValues(infile string) (values []int, err error) {
    file, err := os.Open(infile)
    if err != nil {
        return
    }
    defer file.Close()

    br := bufio.NewReader(file)
    values = make([]int, 0)
    for {
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
                break
            }
        }
        if isPrefix {
            fmt.Println("A too long line, seems unexpected")
            break
        }
        if len(line) < 1 {
            break
        }
        str := string(line) // 字节数组转为字符串
        value, err1 := strconv.Atoi(str)
        if err1 != nil {
            err = err1
            break
        }

        values = append(values, value)
    }
    return
}

func writeValues(values []int, outfile string) error  {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("create file failed.", outfile)
        return err
    }
    defer file.Close()
    for _, v := range values {
        str := strconv.Itoa(v)
        file.WriteString(str + "\n")
    }
    return nil
}

func main() {
    flag.Parse()
    if infile != nil {
        fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
    }
    fmt.Println(*ip)

    values, err := readValues(*infile)
    if err != nil {
        fmt.Println(err)
    } else {
        t1 := time.Now()
        switch *algorithm {
            case "qsort":
                qsort.QuickSort(values)
            case "bubblesort":
                bubblesort.BubbleSort(values)
            default:
                fmt.Println("sort algorithm", *algorithm, "unsupport")
        }
        t2 := time.Now()
        fmt.Println("this process costs", t2.Sub(t1), "to complete")
        writeValues(values, *outfile)
    }
}
