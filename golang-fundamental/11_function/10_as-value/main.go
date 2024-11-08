package main

import "fmt"

func main() {
	// mendeklarasikan ke sebuhah variable
	var sum func(int, int) int = kali
	fmt.Println(sum(2, 2))
}

func kali(param1, param2 int) int {
	return param1 * param2
}
