package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	data := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increase(&data)
	}
	wg.Wait()
	fmt.Println(data)
}

func increase(num *int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutex.Lock()
		*num++
		mutex.Unlock()
	}
}
