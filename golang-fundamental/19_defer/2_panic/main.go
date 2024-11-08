package main

import "fmt"

func main() {
	x := 5
	y := 0

	defer fmt.Println("SESUDAH ========")
	fmt.Println("SEBELUM ========")
	fmt.Println(x / y)

}
