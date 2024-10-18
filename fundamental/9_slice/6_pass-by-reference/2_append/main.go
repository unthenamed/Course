package main

import "fmt"

func main() {

	nameSlice := []int{5, 5, 5, 5}
	appendSlice := append(nameSlice, 5)
	fmt.Println(nameSlice)
	fmt.Println(appendSlice)
	fmt.Println()

	nameSlice[0] = 1
	appendSlice[4] = 1

	fmt.Println(nameSlice)
	fmt.Println(appendSlice)
	fmt.Println()

	/* slice yang di-append adalah salinan dari slice asli,
	   sehingga perubahan pada salah satu slice
	  tidak langsung memengaruhi yang lain(pass by value) */
}
