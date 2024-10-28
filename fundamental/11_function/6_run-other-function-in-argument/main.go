package main

import "fmt"

func main() {
	argument1 := 2
	argument2 := 2

	// 2 + (2x2) = 6
	result := penjumlahan(argument1, perkalian(argument1, argument2))
	fmt.Println("hasil perhitungan :", result)
}

func perkalian(parameter1, parameter2 int) int {
	result := parameter1 * parameter2
	return result
}
func penjumlahan(parameter1, parameter2 int) int {
	result := parameter1 + parameter2
	return result
}
