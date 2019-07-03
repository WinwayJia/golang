package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// use sync.WaitGroup
func demo1() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("done 1")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("done 2")
		defer wg.Done()
	}()

	wg.Wait()
}

// use chan to send signal
func demo2() {
	sig := make(chan bool)
	go func() {
		for {
			select {
			case <-sig:
				fmt.Println("done")
				return
			default:
				fmt.Println("default")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	sig <- true
}

// use context
func demo3() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("default")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func demo4() {
	ctx, cancel := context.WithCancel(context.Background())
	watch := func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println(name, " watching")
				time.Sleep(time.Second)
			}
		}
	}

	go watch(ctx, "wacth1")
	go watch(ctx, "wacth2")
	go watch(ctx, "wacth3")

	time.Sleep(time.Second * 5)
	fmt.Println("stop watching")
	cancel()
	time.Sleep(time.Second * 3)
}

func main() {
	fmt.Println("demo1....................................................................................................")
	demo1()

	fmt.Println("demo2....................................................................................................")
	demo2()

	fmt.Println("demo3....................................................................................................")
	demo3()

	fmt.Println("demo4....................................................................................................")
	demo4()
}
