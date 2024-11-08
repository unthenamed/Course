package main

import "fmt"

func main() {
	argumentData := []int{2, 6, 9, 71, 6, 1, 26}

	kecil, besar := minAndMax(argumentData)
	fmt.Println("terkecil :", kecil, "terbesar", besar)
}

func minAndMax(parameter []int) (int, int) {
	terkecil := parameter[0]
	terbesar := parameter[0]

	for _, value := range parameter {

		// jika value kurang dariterkecil maka rubahlah terkecil
		if value < terkecil {
			terkecil = value
		}

		// jika value lebih besar dari terbesar maka rubahlah terbesarsar
		if value > terbesar {
			terbesar = value
		}

	}

	return terkecil, terbesar
}
