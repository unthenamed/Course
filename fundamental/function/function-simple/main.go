package main

import "fmt"

var name = "abdul" // scope global
func main() { //entry point
	fmt.Println(name)
	helloWorld()
	greet("khoironi", 21)
	add(7, 2)
}

func helloWorld() {
	var name = "jalil" // scope local
	fmt.Println("Hello World", name)
}

func greet(name string, age int) {
	fmt.Println(name, age)
}

func add(a int, b int) {
	fmt.Println(a + b)
}
