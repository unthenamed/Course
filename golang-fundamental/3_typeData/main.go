package main

import "fmt"

func main() {
	fmt.Println("Tipe Data\n")

	fmt.Println("1). Numerik")
	fmt.Println("\tBilangan Bulat")
	// uint khusus menampung bilangan positif
	var bilanganPositif uint8 = 90
	fmt.Printf("\tBilangan positif : %d\n", bilanganPositif)

	// int untuk menampung bilangan positif dan negatif
	var bilanganNegatif int8 = -90
	var intPositif int8 = 90
	fmt.Printf("\tBilangan negatif : %d\n", bilanganNegatif)
	fmt.Printf("\tBilangan positif : %d\n\n", intPositif)

	// float untuk menampung bilangan desimal atau pecahan
	fmt.Println("\tBilangan Pecahan")
	var bilanganPecahan float64 = 3.513
	fmt.Printf("\tBilangan Pecahan : %.2f\n\n", bilanganPecahan)

	fmt.Println("2). Boolean")
	// bool di gunakan untuk menapung nilai benar atau salah
	var exist bool = true
	fmt.Printf("\texist ? %t\n\n", exist)

	fmt.Println("3). String")
	var pesan string = "hallo"
	fmt.Printf("\tPesan : %s\n", pesan)
	fmt.Println(len(pesan))

}
