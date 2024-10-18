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
	fmt.Println("len :", len(nameSlice))
	fmt.Println("cap :", cap(nameSlice))

	// jika melakukan penambahan di luar kapasitas
	/*
		nameSlice[7] = "weekend"
		fmt.Println(nameSlice)
	*/
	// maka akan error [index out of range [7] with length 7]
}
