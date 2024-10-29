package main

import "fmt"

type habitat struct {
	lingkungan string
	suhu       string
}

type hewan struct {
	pakan string
	berat int
	habitat
}

func main() {
	kambing := hewan{
		pakan: "rumput",
		berat: 50,
		habitat: habitat{
			lingkungan: "darat",
			suhu:       "sedang",
		},
	}

	fmt.Println(kambing)
	fmt.Println(kambing.habitat.suhu)
}
