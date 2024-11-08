package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go counter("kelinci", 500, &wg)
	go counter("ayam", 500, &wg)
	wg.Wait()
	fmt.Println("Done")
}

func counter(label string, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i, label)
		time.Sleep(time.Millisecond * duration)
	}
}
