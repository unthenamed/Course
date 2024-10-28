package main

import "fmt"

func main() {
	nama := "abdul"
	alamat := "Kp Nusa "

	fmt.Println("Alamat Memori", nama, "=", &nama)
	fmt.Println("Alamat Memori", alamat, "=", &alamat)
}
