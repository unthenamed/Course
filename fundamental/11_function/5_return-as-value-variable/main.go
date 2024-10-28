package main

import "fmt"

func main() {
	argument1 := 98
	argument2 := 27

	result := test(argument1, argument2)
	fmt.Println(result)
}
func test(parameter1, parameter2 int) int {
	result := parameter1 + parameter2

	// tipe data dari value yang akan di return
	// harus sesuai dengan tipedata yang di deklarasikan
	return result
}
