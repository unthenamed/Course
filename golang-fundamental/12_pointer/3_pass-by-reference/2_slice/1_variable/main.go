package main

import "fmt"

func main() {
	myNumber := []int{8, 2, 3, 1, 6, 7, 4}
	myAnotherNumber := myNumber

	myAnotherNumber[0] = 9

	fmt.Println(myNumber) // ikut berubah
	fmt.Println(myAnotherNumber)

}
