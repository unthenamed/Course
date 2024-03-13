package main

import (
	"fmt"
)

func main() {

	// Anonymous Function
	func() {
		fmt.Println("Hello World")
	}()

	// Anonymous Function dalam variable
	hallo := func() {
		fmt.Println("hallo dunia")
	}
	hallo()

	//Passing Argument ke dalam Anonymous Function
	func(world string) {
		fmt.Println("hallo", world)
	}("jalil")

	// Passing argument dalam variable function
	halo := func(world string) {
		fmt.Println("hallo", world)
	}
	halo("khoironi")

	//passing anonymous function sebagai argument
	english := func(nama string) string {
		return "hello " + nama
	}
	indonesia := func(nama string) string {
		return "halo " + nama
	}
	greet("abdul", english)
	greet("jalil", indonesia)

	tambah := func(x int, y int) int {
		return x + y
	}
	bagi := func(x int, y int) int {
		return x / y
	}
	kurangi := func(x int, y int) int {
		return x - y
	}
	kali := func(x int, y int) int {
		return x * y
	}

	calculate(5, 5, tambah)
	calculate(5, 5, bagi)
	calculate(5, 5, kurangi)
	calculate(5, 5, kali)
}

func greet(nama string, f func(string) string) {
	fmt.Println(f(nama))
}

func calculate(a int, b int, operator func(x int, y int) int) {
	fmt.Println(operator(a, b))
}
