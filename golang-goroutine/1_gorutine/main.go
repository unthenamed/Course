package main

import (
	"fmt"
	"time"
)

func main() {
	go count("hiu")
	fmt.Scanln()
	fmt.Println("selesai")
}

func count(label string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, label)
		time.Sleep(time.Millisecond * 500)
	}
}
