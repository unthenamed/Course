package main

import "fmt"

func main() {
	argument1 := "abdul"
	argument2 := "jalil"
	argument3 := 22

	test(argument1, argument2)
	test1(argument1, argument3)
}

func test(parameter1, parameter2 string) {
	fmt.Println(parameter1, parameter2)
}

func test1(parameter1 string, parameter2 int) {
	fmt.Println("Nama :", parameter1, "Usia :", parameter2)
}
