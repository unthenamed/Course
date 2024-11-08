package main

import "fmt"

func main() {
	defer tambah(2, kali(2, 2))
	fmt.Println("START")
}

func tambah(p1, p2 int) {
	fmt.Println("tambah", p1+p2)
}

func kali(p1, p2 int) (result int) {
	result = p1 * p2
	fmt.Println("kali ", result)
	return
}
