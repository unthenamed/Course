package main

import "fmt"

func main() {

	var nilaiUjian int = 70
	var absen string = "C"

	nilaiAbsen(absen, &nilaiUjian)
	fmt.Println(nilaiUjian)
}

func nilaiAbsen(absensi string, nilai *int) {
	if absensi == "A" {
		*nilai += 5

	}

	if absensi == "B" {
		*nilai -= 5

	}

	if absensi == "C" {
		*nilai -= 10

	}
}
