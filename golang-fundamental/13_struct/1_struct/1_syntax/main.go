package main

import "fmt"

type namaStruct struct {
	nama   string
	umur   int
	alamat string
}

func main() {
	var data1 namaStruct

	fmt.Println(data1) //struct kosong
}
