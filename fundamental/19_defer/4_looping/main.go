package main

import "fmt"

func main() {
	defer fmt.Println("START")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
