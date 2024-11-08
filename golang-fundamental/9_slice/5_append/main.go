package main

import "fmt"

func main() {

	nameSlice := make([]int, 7)

	fmt.Println(nameSlice)
	fmt.Println("len :", len(nameSlice))
	fmt.Println("cap :", cap(nameSlice))

	nameSlice = append(nameSlice, 1, 2)
	fmt.Println(nameSlice)
	fmt.Println("len :", len(nameSlice))
	fmt.Println("cap :", cap(nameSlice))
	// jika kapasitas sudah habis maka jumlah kapaaitas akan di kalikan 2 ( contoh, daei 7 .mainenjadi 14 )
}
