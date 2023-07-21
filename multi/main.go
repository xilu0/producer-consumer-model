package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	// 多个生产者
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				ch <- id + j
				wg.Add(1)
			}
		}(i)
	}

	// 多个消费者
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				wg.Done()
				v, ok := <-ch
				if !ok {
					// channel已经关闭
					return
				}
				fmt.Printf("Consumer %d received: %d\n", id, v)
			}
		}(i)
	}

	// 等待所有任务完成
	wg.Wait()
	// 关闭channel
	close(ch)
}
