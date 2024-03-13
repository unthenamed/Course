package main

import (
	"fmt"
)

func main() {
	num := []int{7, 16, 18, 2, 16, 19, 15}

	evenNumbers := filter(num, isEven)
	tanpaKurung(evenNumbers, "Ganjil :")

	oddNumbers := filter(num, isOdd)
	tanpaKurung(oddNumbers, "Genap :")
}

func filter(numbers []int, f func(int) bool) (result []int) {

	for _, v := range numbers {
		if f(v) {
			result = append(result, v)
		}
	}
	return
}

func isEven(numbers int) bool {
	return numbers%2 == 0

}

func isOdd(numbers int) bool {
	return numbers%2 != 0

}

func tanpaKurung(numbers []int, judul string) {
	fmt.Printf("%s", judul)
	for _, v := range numbers {
		fmt.Printf(" %d", v)
	}
	fmt.Printf("\n")
}
