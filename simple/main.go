package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println("Received", i)
		}
	}()

	wg.Wait()
}
