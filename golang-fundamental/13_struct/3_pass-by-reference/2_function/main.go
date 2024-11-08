package main

import "fmt"

type ayam struct {
	berat int
	harga int
}

func main() {
	kampung := ayam{3, 23000}
	updateHarga(&kampung)

	fmt.Println("main :", kampung) //berubah
}

func updateHarga(param *ayam) {
	param.harga += 2000

	fmt.Println("function :", *param)
}
