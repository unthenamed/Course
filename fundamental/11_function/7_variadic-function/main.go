package main

import "fmt"

func main() {
	result := sum(2, 5, 7, 6, 9, 5, 7)
	fmt.Println(result)

}

func sum(parameter ...int) int {
	fmt.Printf("Tipe data dari parameter %T\n", parameter)
	// sudah di ketahui tipe data parameter adalah slice

	// karna ini slice kurangi satu agar sama seperti indexnya
	length := len(parameter) - 1

	var total int
	//kita lakukan forrange
	for index, value := range parameter {
		total += value

		// tampilan kondisi di terminal
		if index == length {
			fmt.Printf("%d = ", value)

		} else {
			fmt.Printf("%d + ", value)
		}

	}
	return total
}
