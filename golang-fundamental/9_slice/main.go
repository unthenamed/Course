package main

import "fmt"

func main() {
	numberSlice := []int{2, 9, 5, 6} // format dasar slice
	fmt.Println(numberSlice)
	fmt.Println(len(numberSlice)) // tampilkan panjang
	fmt.Println(cap(numberSlice)) // tampilkan kapasitas

	score := make([]int, 4, 7) //membuat slice menggunakan make dengar variable panjang dan kapasitas (optional)
	score[0] = 12345678
	score[1] = 2
	score[2] = 4
	score[3] = 3
	/* score[4] = 7 */ //error
	fmt.Println(score)
	fmt.Println(len(score))
	fmt.Println(cap(score))

	score2 := make([]int, 4) // membuat slice tanpa variable kapasitas
	score2[0] = 2
	score2[1] = 2
	score2[2] = 2
	score2[3] = 2
	fmt.Println(score2)
	fmt.Println(len(score2))
	fmt.Println(cap(score2))

	/* jika array tidak bisa di tambhkan element berbeda dengan slice */

	score3 := [4]string{"satu", "dua", "tiga", "empat"} // contoh aray tidak bisa d tambah element
	score3[1] = "2"
	score3[2] = "2"
	score3[3] = "2"
	/* score3[4] = "2" */ // out of bounds
	fmt.Println(score3)
	fmt.Println(len(score3))
	fmt.Println(cap(score3))

	score4 := make([]int, 4) // contoh slice bisa d tambah element
	score4[0] = 2
	score4[1] = 2
	score4[2] = 2
	score4[3] = 2
	fmt.Println(score4)
	fmt.Println(len(score4))
	fmt.Println(cap(score4))
	score4 = append(score4, 5, 12) //tambahkan element dengan append
	fmt.Println(score4)
	fmt.Println(len(score4))
	fmt.Println(cap(score4))

	/* Array itu Pass by Value dan Slice itu Pass by Reference */

	//array
	var score5 = [4]string{"satu", "dua", "tiga", "empat"}
	var otherScore5 = score5 // dalam array adala penduplikasian, variable score5 tidakakan berubah
	fmt.Println(score5)
	fmt.Println(otherScore5)
	otherScore5[0] = "kosong"
	otherScore5[1] = "kosong"
	otherScore5[2] = "kosong"
	otherScore5[3] = "kosong"
	fmt.Println(score5) // masi sama
	fmt.Println(otherScore5)

	/*output
	[satu dua tiga empat]
	[satu dua tiga empat]
	[satu dua tiga empat]
	[kosong kosong kosong kosong]
	*/

	//slice
	var score6 = []int{10, 20, 30}
	var otherScore6 = score6 // dalam slice hanya syntax saja
	fmt.Println(score6)
	fmt.Println(otherScore6)
	otherScore6[0] = 100
	otherScore6[1] = 100
	otherScore6[2] = 100
	fmt.Println(score6) // berubah
	fmt.Println(otherScore6)

	/* output
	[10 20 30]
	[10 20 30]
	[100 100 100]
	[100 100 100]
	*/

	/* membuat slice dari array yang sudah ada */
	var exArray = [4]int{1, 2, 3, 4}
	var sliceExArray = exArray[0:3]
	fmt.Println(exArray)
	fmt.Println(sliceExArray)
	sliceExArray[2] = 12
	fmt.Println(exArray)
	fmt.Println(sliceExArray)
	sliceExArray = append(sliceExArray, 20)
	fmt.Println(exArray)
	fmt.Println(sliceExArray)

	/* output
	[1 2 3 4]			// array
	[1 2 3]				// slice dari array[0:3]
	[1 2 12 4]			// index[2] pada array ikut berubah
	[1 2 12]			// ketika index[2] pada slice berubah
	[1 2 12 20]			// index[3] pada array ikut berubah
	[1 2 12 20]			// ketika variable sliceExArray di Append untuk menambah index[3]
	*/

}
