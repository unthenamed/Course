package main

import "fmt"

func main() {
	tambah := func(param1, param2 int) int {
		return param1 + param2
	}

	kurang := func(param1, param2 int) int {
		return param1 - param2
	}

	calculate(2, 3, tambah)
	calculate(5, 3, kurang)
}

func calculate(param1, param2 int, f func(int, int) int) {
	fmt.Println(f(param1, param2))
}
