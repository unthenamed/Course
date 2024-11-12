package main

import (
	"fmt"
)

var channel = make(chan int)

func rataRata() {
	var jumlah int
	var penjumlah int
	for nilai := range channel {
		jumlah += nilai
		penjumlah++

		fmt.Println("rata-rata :", jumlah/penjumlah)
	}
}

func insert(v ...int) {
	for _, n := range v {
		channel <- n
	}
	close(channel)
}

func main() {
	go insert(57, 42, 89, 21, 617, 627, 156, 1628)
	rataRata()
}
