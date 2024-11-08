package main

import "fmt"

func main() {
	myNumber := []int{8, 2, 3, 1, 6, 7, 4}

	sortGanjil(&myNumber)
	fmt.Println(myNumber)
}

func sortGanjil(param *[]int) {
	result := make([]int, 0)

	for _, value := range *param {
		if value%2 != 0 {
			result = append(result, value)
		}
	}

	*param = result
}
