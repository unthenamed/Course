package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// inisialisasi scanner bufio
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Register Akun Unthenamed")

	// inputan untuk string
	fmt.Print("\tNama : ")
	scanner.Scan()
	userName := scanner.Text()

	// inputan untuk int menggunakan strconv
	fmt.Print("\tUmur : ")
	scanner.Scan()
	age, _ := strconv.Atoi(scanner.Text())
	age = (2024 - age)

	fmt.Printf("Nama anda adalah %s, \ndan anda lahir di tahun %d \n", userName, age)

}
