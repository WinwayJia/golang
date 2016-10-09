package main
import (
	"time"
	"fmt"
	"math/rand"
)

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // 将计算结果发送到channel中
}

func main() {
//	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	resultChan := make(chan int, 2)
//	go sum(values[:len(values)/2], resultChan)
//	go sum(values[len(values)/2:], resultChan)
//	sum1, sum2 := <-resultChan, <-resultChan // 接收结果
//	fmt.Println("Result:", sum1, sum2, sum1 + sum2)
	
	// 
	inChan := make(chan []int, 1024)
	outChan := make(chan int64)
	
	
	data := make([]int, 32)
	for i:=0; i<32; i+=8 {
		go gen_random(data[i:i+8])
		inChan <- data[i:i+8]
		
	}
	
	go sum_random(inChan, outChan)
	go sum_random(inChan, outChan)
	go sum_random(inChan, outChan)
	go sum_random(inChan, outChan)
	
	sum1, sum2, sum3, sum4 := <-outChan, <-outChan, <-outChan, <-outChan
	fmt.Println(sum1, sum2, sum3, sum4)
	fmt.Printf("sum: %d\n", sum1 + sum2 + sum3 + sum4)
	
	test_chan_time()
}

func gen_random(data []int) {
	for i:=0; i<len(data); i++ {
		data[i] = rand.Intn(64)
	}
}

func sum_random(inChan chan []int, outChan chan int64) {
	var result int64 = 0
	for {
		val := <- inChan
		for _, v := range val {
			result += int64(v)
		}
		outChan <- result
	}
	
}


func test_chan_time() {
	chanTimeout := make(chan bool)
	chanInt := make(chan int)
	go func () {
		time.Sleep(time.Second*3)
		chanTimeout <- true
	}()
	select {
		case <-chanTimeout:
			fmt.Println("Timeout")
		case <-chanInt:
//		default:
//			fmt.Println("default")
	}
}