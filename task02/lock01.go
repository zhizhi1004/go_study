package main

import (
	"fmt"
	"sync"
)

var sum int = 0
var mu sync.Mutex

func add() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		sum++
		mu.Unlock()
	}

}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			add()
		}()
	}
	wg.Wait()
	fmt.Printf("合计:%d\n", sum)
}
