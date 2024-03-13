package main

import "fmt"

func main() {

	num := []int{2, 6, 9, 12, 6}
	kecil, besar := minMax(num)
	fmt.Println("kecil", kecil)
	fmt.Println("besar", besar)

	_, kecil = maxMin(num)
	fmt.Println("kecil", kecil)

}

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

//Named return value
func maxMin(numbers []int) (max int, min int) {
	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}
