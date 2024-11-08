package main

import "fmt"

func main() {
	argData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bilanganGanjil := cari(argData, ganjil)
	bilanganGenap := cari(argData, genap)

	fmt.Println(bilanganGanjil)
	fmt.Println(bilanganGenap)

}

func cari(param []int, fungsi func(int) bool) (result []int) {
	for _, value := range param {
		if fungsi(value) {
			result = append(result, value)
		}
	}
	return
}

func ganjil(param int) bool {
	return param%2 == 0
}

func genap(param int) bool {
	return param%2 != 0
}
