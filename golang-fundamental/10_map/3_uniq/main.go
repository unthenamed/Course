package main

import "fmt"

func main() {
	user := make(map[string]int)
	fmt.Println(user)

	fmt.Println(user["total"])

	user["total"] = 12
	user["aktif"] = 5
	fmt.Println(user)

	user["aktif"] = 7
	fmt.Println(user)
}
