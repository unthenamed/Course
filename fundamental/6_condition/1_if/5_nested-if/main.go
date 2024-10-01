package main

import "fmt"

func main() {
	lebar := 70
	tinggi := 60
	minimum := 50

	if lebar >= minimum {
		if tinggi <= minimum {
			fmt.Printf("lebar %d lebih besar dan tinggi %d lebih kecil dari minimum %d \n", lebar, tinggi, minimum)
		} else {
			fmt.Printf("lebar %d lebih besar dari minimum %d \n ", lebar, minimum)
		}

	}
}
