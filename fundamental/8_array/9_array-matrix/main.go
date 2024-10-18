package main

import "fmt"

func main() {

	var lock = [...]int{2, 5, 8, 9, 1, 4, 6, 7}

	var pattern = [3][3]int{}
	for i, v := range lock {
		v -= 1
		i += 1
		var index int
		var matrix int

		if v <= 2 {
			index = 0
			if v == 0 {
				matrix = 0
			} else if v == 1 {
				matrix = 1
			} else if v == 2 {
				matrix = 2
			}
		} else if v <= 5 {
			index = 1
			if v == 3 {
				matrix = 0
			} else if v == 4 {
				matrix = 1
			} else if v == 5 {
				matrix = 2
			}
		} else if v <= 8 {
			index = 2
			if v == 6 {
				matrix = 0
			} else if v == 7 {
				matrix = 1
			} else if v == 8 {
				matrix = 2
			}
		}

		pattern[index][matrix] = i
	}

	for n := range pattern {
		for i := 0; i < len(pattern[n]); i++ {
			fmt.Printf(" %d ", pattern[n][i])
		}
		fmt.Println()
	}
}
