package main

import "fmt"

type kursi struct {
	panjang int
	lebar   int
	harga   int
}

func main() {

	pantai := new(kursi)
	fmt.Println(pantai)
	fmt.Printf("%p", pantai)

	// pantai menjadi variable bertipee pointer
}
