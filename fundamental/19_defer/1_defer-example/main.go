package main

import "fmt"

func main() {
	x := 4
	defer fmt.Println(x)

	x += 4
	fmt.Println(x)
}
