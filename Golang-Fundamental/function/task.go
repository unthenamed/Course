/*

Setiap bambu memiliki sekat.
Buatlah aplikasi yang dapat menentukan jumlah bambu yang kamu miliki melalui inputan,
setelah itu kamu dapat menentukan jumlah sekat dari setiap bambu yang kamu miliki.
Jika sudah, potonglah bambu beberapa kali berdasarkan inputan yang kamu inginkan.
Setiap potongan yang kamu lakukan akan menghilangkan 1 sekat disetiap bambu yang kamu miliki.
Lalu tampilkanlah pada console sisa sekat bambu setiap pemotongan.


Kriteria:
  > Gunakan variabel untuk menerima inputan dari console
  > Gunakan function

Sample Input:
Masukan banyak bamboo : 2
sekat bamboo ke-1 : 5
sekat bamboo ke-2 : 3
Berapa kali potong? 2



Sample Output :

Potongan ke- 1
4
2

Potongan ke- 2
3
1
*/
package main

import (
	"fmt"
)

func main() {
	var banyakBambu int
	var potonganBambu int

	fmt.Scanln(&banyakBambu)
	sekatBambu := input(banyakBambu)
	fmt.Scanln(&potonganBambu)
	jumlahPotongan(potonganBambu, sekatBambu)

}

func input(banyakBambu int) (result []int) {
	for i := 0; i < banyakBambu; i++ {
		var input int
		fmt.Scanln(&input)
		result = append(result, input)
	}
	return
}

func potongan(number []int) {
	for i, v := range number {
		if v != 0 {
			number[i] = v - 1
		}
		fmt.Println(number[i])
	}
}

func jumlahPotongan(jumlah int, sekat []int) {
	for i := 0; i < jumlah; i++ {
		fmt.Println("Potongan ke-", i+1)
		potongan(sekat)
	}
}
