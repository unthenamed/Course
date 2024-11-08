package main

import "fmt"

func main() {

	myArray := [2]string{"satu", "dua"}

	change(myArray)
	fmt.Println(myArray)

}

func change(param [2]string) {
	for index, value := range param {
		param[index] = "hallo " + value
	}

	fmt.Println(param)
}
