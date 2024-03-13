package main

import (
	"fmt"
)

func main() {

	result := "lulus"
	if result == "lulus" {
		fmt.Println("anda", result)
	} else {

	}

	if hasil := 25; hasil == 30 {
		fmt.Println("hasil", hasil)
	} else {

	}

	gpa := 3.0
	var status string
	if gpa >= 3.5 && gpa <= 4.0 {
		status = "memuaskan"
	} else if gpa >= 3.0 && gpa <= 3.5 {
		status = "oke lah"
	} else if gpa >= 2.75 && gpa <= 3.0 {
		status = "mantap"
	} else {
		status = "tidak memenuhi syarat"
	}
	fmt.Printf(" lulus dengan predikat %s dengan Gpa %.2f\n", status, gpa)

	x := 25
	y := 30
	if x > 0 {
		if y > 0 {
			fmt.Println("x dan y > 0")
		} else {
			fmt.Println("x > 0 dan y < 0")
		}
	}

	jam := 5
	switch jam {
	case 8, 9:
		{
			fmt.Println("8-9")
		}
		fallthrough
	case 10:
		{
			fmt.Println("10")
		}

	}

	jam = 1
	if jam <= 12 {
		switch {
		case jam >= 4 && jam <= 5:
			{
				fmt.Println("Bangun Pagi")
			}
		case jam >= 6 && jam <= 7:
			{
				fmt.Println("Mandi dan Sarapan")
			}
		case jam >= 8 && jam <= 11:
			{
				fmt.Println("Berangkat Sekolah")
			}
		case jam == 12:
			{
				fmt.Println("Pulang Sekolah")
			}
		case jam >= 13 && jam <= 15:
			{
				fmt.Println("Tidur Siang")
			}
		case jam >= 16 && jam <= 17:
			{
				fmt.Println("Waktu Main")
			}
		default:
			{
				fmt.Println("Waktu Belajar dan Istirahat")
			}
		}
	} else {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	}

}
