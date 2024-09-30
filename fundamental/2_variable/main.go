package main

import "fmt"

func main() {

	// Manifest typing
	// (tipe data di tuliskan atau di tentukan )
	var firstName string

	// Mengisi nilai variable
	// (jika Manifest typing niali harus sesuai
	// dengan variable yang di deklarasikan)
	firstName = "Abdul"

	// namun lebih baik jika nilai di tuliskan
	// pada saat mendeklarasikan Variable
	var lastName string = "Jalil"

	//	Atau bisa juga Mendeklarasikan
	//	banyak variable sekaligus
	var (
		age     int
		address string
	)
	age = 22
	address = "Kp Nusa"

	// Atau bisa juga seperti ini
	var (
		namaSekolah, alamatSekolah = "SMU", "Jawilan"
	)

	// print value Variable ke terminal
	fmt.Println(firstName, lastName)
	fmt.Println(namaSekolah, alamatSekolah)
	fmt.Printf("hallo, nama saya %s %s, umur saya %d saya tinggal di %s \n", firstName, lastName, age, address)

	//	type interest
	//	Mendeklarasikan variable sesuai dengan
	//	tipedata yang di berikan
	hari := "Senin"
	tanggal := 25
	bulan := "Agustus"
	fmt.Println(hari, tanggal, bulan+" 2024")

	// Constant
	// const adalah sebuah variable yang tidak bisa
	// di rubah nilainya
	const phi = 3.14
	fmt.Println(phi)
	// jika di rubah
	// phi = 2.14
	// maka akan terjadi error

	// Variable underscore
	// digunakan untuk menampung nilai tidak di pakai
	pekerjaan, _ := "karyawan", "Operator"
	fmt.Println(pekerjaan)

}
