package main
import "math/rand"
import "os"
import "strconv"

func TestBubbleSort2(outfile string) (err error) {
    file, err := os.Create(outfile)
    if err != nil {
        return
    }

    for i:=0; i<10000; i++ {
        str := strconv.Itoa(rand.Intn(65536))
        file.WriteString(str + "\n")
    }

    return nil
}

func main() {
    TestBubbleSort2("unsorted.dat")
}
