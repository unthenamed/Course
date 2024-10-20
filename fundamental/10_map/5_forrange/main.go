package main

import "fmt"

func main() {
	items := map[int]string{
		1: "kue",
		2: "air",
		3: "rokok",
	}
	fmt.Println(items)

	for key, value := range items {
		fmt.Println(key, value)
	}
}
