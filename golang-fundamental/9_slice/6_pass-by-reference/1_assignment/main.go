package main

import "fmt"

func main() {

	nameSlice := []int{5, 5, 5, 5}
	pointerSlice := nameSlice
	fmt.Println(nameSlice)
	fmt.Println(pointerSlice)
	fmt.Println()

	nameSlice[0] = 1
	pointerSlice[3] = 1
	fmt.Println(nameSlice)
	fmt.Println(pointerSlice)
	fmt.Println()
}

/* keduanya merujuk ke array yang sama. Perubahan pada salah satu slice akan merubah slice yang lain.*/
