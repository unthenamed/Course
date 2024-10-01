package main

import "fmt"

func main() {

	nilaiUjian := 71

	predikatA := 90
	predikatB := 80

	var hasilUjian string

	if nilaiUjian >= predikatA && nilaiUjian <= 100 {
		hasilUjian = "Terbaik"
	} else if nilaiUjian >= predikatB && nilaiUjian < predikatA {
		hasilUjian = "Baik"
	} else if nilaiUjian >= 70 && nilaiUjian < predikatB {
		hasilUjian = "Masuk"
	} else {
		hasilUjian = "Gagal"
	}

	fmt.Println(hasilUjian)
}
