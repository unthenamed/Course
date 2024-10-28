package main

import "fmt"

func main() {

	myArray := [2]string{"satu", "dua"}
	mySecArray := myArray

	fmt.Println("Nilai Awal")
	fmt.Println(myArray)
	fmt.Println(mySecArray)

	myArray[0] = "diubah"

	fmt.Println("Nilai sesudah")
	fmt.Println(myArray)
	fmt.Println(mySecArray)

}
