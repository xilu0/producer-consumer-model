package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	// 创建一个带有取消功能的context
	ctx, cancel := context.WithCancel(context.Background())

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-ch:
				fmt.Println("Received", v)
				if v > 5 {
					// 当接收到的值大于5时，取消所有的goroutine
					cancel()
				}
			}
		}
	}()

	wg.Wait()
}
