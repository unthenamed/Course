package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	channelInt := make(chan int, 5)
	wg.Add(1)
	send(channelInt, 5)
	send(channelInt, 5)
	send(channelInt, 5)
	send(channelInt, 5)
	send(channelInt, 5)
	go get(channelInt)

	wg.Wait()
}

func send(c chan int, num int) {
	c <- num
	fmt.Println("channel <-", num)

}

func get(c chan int) {
	defer wg.Done()

	for len(c) > 0 {
		fmt.Println("\t\t", <-c, "<- channel")
		time.Sleep(time.Second * 1)
	}
}
