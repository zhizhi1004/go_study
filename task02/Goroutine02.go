package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	task := []struct {
		name     string
		duration time.Duration
	}{
		{"task1", time.Second * 1},
		{"task2", time.Second * 2},
		{"task2", time.Second * 3},
	}

	wg.Add(len(task))
	for _, t := range task {
		go runTask(t.name, t.duration, &wg)
	}
	wg.Wait()

}

func runTask(name string, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	time.Sleep(duration)
	costTime := time.Since(startTime)
	fmt.Println(name, " cost : ", costTime)
}
