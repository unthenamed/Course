package main

import "fmt"

func main() {
	fmt.Println("Logika")

	a := true
	b := false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)

	// && akan di jalankan terlebuih dahulu
	fmt.Println(false || true && false)

}
