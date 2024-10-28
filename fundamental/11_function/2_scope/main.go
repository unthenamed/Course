package main

import "fmt"

// Scope global
var name string = "jalil"
var address string = "Kp Nusa"

func main() {
	// Scope Local
	name := "Abdul"
	test()
	fmt.Println(name, address)
}

func test() {
	fmt.Println(name, address)
}
