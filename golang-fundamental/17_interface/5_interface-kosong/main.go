package main

import "fmt"

func main() {

	//deklarasi interface kosong
	var data interface{}
	var data1 any

	// ketika di assigment tipe data akan berubah
	// menyesuaikan value yang di inputkan

	data = "jalil"
	data = 0
	data = true

	data1 = "jalil"

	fmt.Println(data)
	fmt.Println(data1)

}
