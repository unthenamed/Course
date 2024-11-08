package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Print(" ", i)

		if i == 10 {
			break
		}

		i++

	}
}
