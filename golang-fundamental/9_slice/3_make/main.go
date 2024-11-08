package main

import "fmt"

func main() {

	nameSlice := make([]int, 7, 14)

	fmt.Println(nameSlice)
	fmt.Println("len :", len(nameSlice))
	fmt.Println("cap :", cap(nameSlice))
}
