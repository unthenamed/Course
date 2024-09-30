package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var age int

	// input untuk string
	fmt.Println("Masukan nama anda :")
	fmt.Scanln(&firstName, &lastName)

	// input untuk int
	fmt.Println("Masukan umur anda :")
	fmt.Scanln(&age)
	tahunLahir := (2024 - age)

	fmt.Println("Nama anda : ", firstName, lastName)
	fmt.Println("Anda lahir di tahun : ", tahunLahir)
}
