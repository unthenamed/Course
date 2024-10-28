package main

import "fmt"

func main() {
	argument1 := 12
	argument2 := 15

	fmt.Println("jumlah :", test(argument1, argument2))
}

func test(parameter1, parameter2 int) int {
	return parameter1 + parameter2
}
