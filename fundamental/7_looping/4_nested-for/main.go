package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Print(" ", j)
		}
		fmt.Println()
	}
}
