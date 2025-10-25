package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sun int32 = 0

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&sun, 2)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("sun : %d\n", sun)
}
