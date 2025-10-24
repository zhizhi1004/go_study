package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	//生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Println("生产者:", i)
		}
		fmt.Println("生产者结束")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("消费者:", v)
		}
	}()
	wg.Wait()
	fmt.Println("所有任务结束")

}
