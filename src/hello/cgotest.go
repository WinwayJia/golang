package main

// #cgo LDFLAGS: -lpthread
// #cgo CFLAGS: -Wimplicit-function-declaration
/*
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <pthread.h>
int Print() {
	printf("Hello cgo\n");
}
void* ThreadFunc(void* arg) {
	while (1) {
		Print();
		sleep(3);
	}
}
pthread_t tid;
int StartThread() {
	pthread_create(&tid, NULL, ThreadFunc, NULL);
}

*/
import "C"

import (
	"fmt"
	"time"
)

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(100)
	fmt.Println(Random())

	str := []byte("Helloa")
	fmt.Printf("%v\n", &str)
	//	str += " World"
	fmt.Printf("%v\n", &str)
	str[0] = 'C'
	fmt.Printf("%s", str[:5])

	if err != nil {
		fmt.Println(err)
	}
	C.Print()
	C.StartThread()
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("main thread")
	}
}
