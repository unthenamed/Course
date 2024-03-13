package main

import "fmt"

func main() {

	fmt.Println(add(2, 7))
	total := add(7, 8)
	fmt.Println(total)

	perkalian := multiply(7, 8)
	fmt.Println(perkalian)

	hasilKalkulasi := add(15, multiply(27, 6))
	fmt.Println(hasilKalkulasi)
}

func add(a int, b int) int {
	result := a + b
	return result
}

func multiply(a, b int) int {
	result := a * b
	return result

}
