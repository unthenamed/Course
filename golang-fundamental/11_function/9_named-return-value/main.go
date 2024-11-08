package main

import "fmt"

func main() {
	ganjil, genap := ganjilGenap(2, 4, 5, 8, 7, 6, 16, 17)
	fmt.Println(ganjil)
	fmt.Println(genap)
}

func ganjilGenap(parameter ...int) (ganjil []int, genap []int) {
	for _, value := range parameter {

		if value%2 != 0 {
			ganjil = append(ganjil, value)
		}

		if value%2 == 0 {
			genap = append(genap, value)
		}
	}
	return
}
