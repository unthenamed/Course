package main

import "fmt"

func main() {

	nameSlice := []string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu"}

	fmt.Println(nameSlice)

	nameSlice[6] = "libur"
	fmt.Println(nameSlice)
}
