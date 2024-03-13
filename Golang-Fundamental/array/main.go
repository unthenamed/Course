package main

import "fmt"

func main() {

	var fruit = [4]string{"mntimun", "aple", "semangka", "pir"}
	fruit[0] = "Mangga" // mengubah nilai dari index ke-0 dengan menyimpan string "Mangga"
	fmt.Println(fruit)

	var hari = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "minggu"}
	fmt.Println(len(hari)) // Menampilkan panjang array hari yaitu 7

	for i := 0; i < len(hari); i++ {
		fmt.Printf("%d %s\n", i, hari[i])
	}

	for index, value := range hari {
		fmt.Printf("%d %s\n", index, value)

	}

	for _, value := range hari {
		fmt.Println(value)

	}

	for index := range hari {
		fmt.Println(index)

	}

	var number = [5]int{3, 8, 2, 7, 4}
	for _, numbers := range number {
		if numbers%2 == 0 {
			fmt.Println(numbers)
		}

	}

	fmt.Println(number)
	for i := 0; i < len(number); i++ {
		number[i] *= 2
	}
	fmt.Println(number)

	var matrx = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println("Matrix : ", matrx[2][1])

}
