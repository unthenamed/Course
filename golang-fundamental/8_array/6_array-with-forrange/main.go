package main

import "fmt"

func main() {
	var tahun = [4]int{1990, 1995, 2000, 2005}
	for _, thn := range tahun {
		fmt.Println("tahun :", thn)
	}
}
