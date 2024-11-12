package main

import (
	"fmt"
	"time"
)

func writeString(c chan string, n ...string) {
	for _, value := range n {
		c <- value
	}
	close(c)
}

func main() {
	binatangChannel := make(chan string)
	buahChannel := make(chan string)

	go writeString(binatangChannel, "hiu", "kuda", "sapi", "ayam")
	go writeString(buahChannel, "naga", "limun", "sukun", "pepaya")

	time.Sleep(time.Second * 2)
	var statusA bool
	var statusB bool
	for {
		select {
		case data, status := <-binatangChannel:
			statusA = status

			if statusA {
				fmt.Println("Channel binatang :", data)
			}
		case data, status := <-buahChannel:
			statusB = status

			if statusB {
				fmt.Println("Channel buah :", data)
			}

		}

		if !statusA && !statusB {
			break
		}
	}
}
