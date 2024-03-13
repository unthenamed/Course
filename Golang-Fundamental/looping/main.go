package main

import "fmt"

func main() {

	//  init; condition; post statement;
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka", i)
	}

	/* 	// init statement hanya akan di eksekusi sekali saat awal perulangn
	   	// condition dimana di lakukan pengecekan di setiap perulangan (true di lanjutkan, false di hentikan)
	   	// post statment di lakukan sekali di akhir setiap perulangan
	   	// init => condition {true} => blok kode {$} => post statment {$} => condition {false} */

	//perulangan hanya kondisi
	var i = 6
	for i < 0 {
		fmt.Println("%d I will become a go developer", i)
		i--
	}

	//perulangan tanpa argumen sama kali
	var index = 0
	for {
		fmt.Println("enigma", index)
		index++
		if index == 5 {
			break
		}
	}

	//perulangan di dalam perulangan {Naked looping}
	for j := 1; j <= 3; j++ {

		for j := 1; j <= 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	//Break pada perulangn digunakan untuk keluar dari looping
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break //		|break|======
		} //							||
		fmt.Print(i, " ") //			||	output sesudah nomor 5 lanjut ke done
	} //								||
	fmt.Println("done") //		<========

	// continue pada perulangan di gunakan untuk melanjutkan ke perulangn selanjutnya tanpa mengeksekusi kode blok
	for i := 1; i <= 10; i++ { //  <==========
		if i == 5 { //						|| output nomor 5 tidak terprint
			continue //  |continue| ==========
		}
		fmt.Print(i, " ")
	}
	fmt.Println("done")

	//perulangan di dalam perulangan {Naked looping} menggunakan break/continue
	for j := 1; j <= 3; j++ {

		for j := 1; j <= 5; j++ {
			if j == 4 {
				//break
				continue
			}
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}
