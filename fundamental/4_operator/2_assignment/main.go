package main

import "fmt"

func main() {
	fmt.Println("Assignment")
	var a int = 10
	var b int

	// samadengan
	b = a
	fmt.Println("Nilai b :", b)

	// tambah samadengan
	b += a
	fmt.Println("Nilai b :", b)

	// kurang samadengan
	b -= a
	fmt.Println("Nilai b :", b)

	// kali samadengan
	b *= a
	fmt.Println("Nilai b :", b)

	// bagi samadengan
	b /= a
	fmt.Println("Nilai b :", b)

	// modulus samadengan
	b %= a
	fmt.Println("Nilai b :", b)
}
