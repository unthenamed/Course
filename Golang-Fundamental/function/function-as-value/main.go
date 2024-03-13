package main

import "fmt"

func main() {
	var sum func(int, int) int = add
	fmt.Println(sum(25, 50))
}

func add(a, b int) int {
	return a + b
}
