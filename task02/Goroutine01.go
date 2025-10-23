package main

import (
	"fmt"
	"sync"
)

func PrintEventNumber(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go PrintEventNumber("偶数", &wg)
	go func(str string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(str)
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}("奇数", &wg)
	wg.Wait()
}
