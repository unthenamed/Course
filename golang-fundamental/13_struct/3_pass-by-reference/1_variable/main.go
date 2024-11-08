package main

import "fmt"

type pohon struct {
	habitat string
	hasil   string
	tinggi  int
}

func main() {
	bambu := pohon{"tropis", "kayu", 10}
	bambu2 := &bambu

	bambu2.tinggi = 15

	fmt.Println(bambu) // berubah
	fmt.Println(*bambu2)

}
