package main

import "fmt"

func main() {
	pointer1 := 1
	pointer2 := pointer1

	pointer2++

	// di buktikan pointer satu tidak ikut berubah
	fmt.Println(pointer1)
	fmt.Println(pointer2)

	//karna alamat memorinya berbeda
	fmt.Println("Alamat memori pointer1", &pointer1)
	fmt.Println("Alamat memori pointer2", &pointer2)

}
