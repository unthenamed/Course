package main

import "fmt"

// struct
type kendaraan struct {
	pabrikan string
	keluaran int
	harga    int
}

// method
func (param *kendaraan) updateHarga(harga int) {
	param.harga = harga
}

func main() {

	var brio = kendaraan{"honda", 2010, 150}
	brio.updateHarga(10)
	fmt.Println(brio)
}
