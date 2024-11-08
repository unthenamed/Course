/*
Tugas Menukar Variable
Pada tugas ini kamu diminta melengkapi coding berikut
untuk menukar nilai variabel a dan b
lalu tampilkanlah pada console.

Contoh:
a = Camp, b = Enigma
maka nilai a menjadi Enigma dan b menjadi Camp.
*/
package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	// Tuliskan kode disini

	a, b = b, a
	fmt.Println(a, b)
}
