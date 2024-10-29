package main

import "fmt"

type buah struct {
	warna string
	rasa  string
	harga int
}

func main() {
	semangka := buah{"merah", "manis", 10000}
	semangka2 := semangka

	semangka2.rasa = "pait"

	fmt.Println(semangka) //tidak berubah
	fmt.Println(semangka2)

}
