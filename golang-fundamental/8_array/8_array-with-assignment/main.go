package main

import "fmt"

func main() {

	var nilai = [...]int{70, 75, 80, 85, 90}
	fmt.Println(nilai)

	for i := range nilai {
		nilai[i] += 5
	}
	fmt.Println(nilai)
}
