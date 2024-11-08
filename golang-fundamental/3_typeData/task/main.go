/*
Pada tugas kali ini kamu akan diberikan dua variabel
yang pertama dalam bentuk integer
dan yang kedua dalam bentuk float32.
Tulislah code yang akan menampilkan pada console berupa :

 1. Baris pertama menampilkan penjumlahan
    dari dua variabel yang diberikan dalam bentuk integer
 2. Baris kedua menampilkan pengurangan
    dalam bentuk float dengan 2 koma
 3. Baris ketiga menampilkan perkalian
    dalam bentuk integer

Example :

	x = 2
	y= 3.12

Output:

	5
	-1.12
	6
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int
	var y float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ = strconv.ParseFloat(scanner.Text(), 32)

	// Tulis kode disini
	sumToInt := x + int(y)
	minToFloat := float64(x) - y
	mulToInt := x * int(y)
	fmt.Printf("%d\n%.2f\n%d\n", sumToInt, minToFloat, mulToInt)

}
