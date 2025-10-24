package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("send:", i)
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("recv:", v)
		}
	}()

	time.Sleep(time.Second * 1)
}
