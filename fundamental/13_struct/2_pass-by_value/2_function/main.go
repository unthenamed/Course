package main

import "fmt"

type buah struct {
	warna string
	rasa  string
	harga int
}

func main() {
	anggur := buah{"ungu", "sepat", 15000}
	updateHarga(anggur)

	fmt.Println("main :", anggur)
}

func updateHarga(param buah) {
	param.harga += 2000

	fmt.Println("function :", param)
}
