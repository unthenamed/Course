package main

import "fmt"

func main() {
	var array = [5]int{11, 2, 3, 4, 55}
	fmt.Println(array)
	fmt.Println("len :", len(array))
	fmt.Println("cap :", cap(array))
}
