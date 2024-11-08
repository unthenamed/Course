package main

import (
	"fmt"
	"math"
)

// struct
type segitiga struct {
	a float64
	b float64
	c float64
}

// Method
func (s segitiga) getLuas() float64 {
	semi := s.a + s.b + s.c/2
	result := semi * (semi - s.a) * (semi - s.b) * (semi - s.c)
	return math.Sqrt(result)
}

func (s segitiga) getKeliling() float64 {
	return s.a + s.b + s.c
}

// interface
type bidang interface {
	getLuas() float64
	getKeliling() float64
}

func main() {
	// upcasting (di deklarasi dengan tipedata interface)
	var samakaki bidang = segitiga{5, 5, 5}
	luas(samakaki)
	keliling(samakaki)

	// downcasting (mengembalikan ke tipedata struct)
	var dataSamakaki segitiga = samakaki.(segitiga)
	fmt.Println(dataSamakaki.a)
	fmt.Println(dataSamakaki.b)
	fmt.Println(dataSamakaki.c)
}

// fungsi dengan interface sebagai parameter
func luas(l bidang) {
	fmt.Printf("Luas : %.2f\n", l.getLuas())
}

func keliling(l bidang) {
	fmt.Printf("Keliling : %.2f\n", l.getKeliling())
}
