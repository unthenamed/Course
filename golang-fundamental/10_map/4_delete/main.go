package main

import "fmt"

func main() {

	items := map[int]string{
		1: "sukro",
		2: "basreng",
		3: "milkita",
	}
	fmt.Println(items)

	delete(items, 1)
	fmt.Println(items)
}
