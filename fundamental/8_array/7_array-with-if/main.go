package main

import "fmt"

func main() {
	var bilangan = [...]int{2, 3, 5, 6, 7, 5, 8, 96, 5, 6, 44}
	for i := 0; i < len(bilangan); i++ {
		if bilangan[i]%2 != 0 {
			fmt.Println(bilangan[i])
		}
	}
}
