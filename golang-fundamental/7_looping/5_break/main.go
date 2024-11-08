package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(" ", i)
		if i == 5 {
			break
		}
	}
}
