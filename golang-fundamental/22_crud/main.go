package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unthenamed/libs"
)

type Hp struct {
	id      int
	nama    string
	tipe    string
	wa      string
	sandi   string
	keluhan string
	biaya   int
}

var Perbaikan []Hp
var DiAmbil []Hp
var DiKembalikan []Hp

func main() {
	var exit bool
	for !exit {
		libs.ClearTerm()
		fmt.Println(strings.Repeat("=", 50))
		NewPerbaikan()
		fmt.Println(Perbaikan)
		in, _ := libs.InputNumber("Masukan ID :")
		UpdateDiAmbil(in)
		fmt.Println(Perbaikan)
		exit = true
	}
}

func CekIdPerbaikan(f int) bool {
	for _, data := range Perbaikan {
		if data.id == f {
			return true
		}
	}
	return false
}

func NewPerbaikan() {
	var masuk Hp
	masuk.id = rand.Intn(587926)
	masuk.nama, _ = libs.InputText("	Nama Pemilik :")
	masuk.tipe, _ = libs.InputText("	Tipe Hp :")
	masuk.keluhan, _ = libs.InputText("	Keluhan :")
	masuk.wa, _ = libs.InputText("	Wa/No-HP :")
	masuk.sandi, _ = libs.InputText("	Kunci Layar :")
	masuk.biaya, _ = libs.InputNumber("	Biaya :")
	if !CekIdPerbaikan(masuk.id) {
		Perbaikan = append(Perbaikan, masuk)

		return
	}
}

func UpdateDiAmbil(id int) {
	for _, data := range Perbaikan {
		if data.id == id {
			if nama, err := libs.InputText("	Nama Pemilik :"); !err {
				data.nama = nama
			}
			if tipe, err := libs.InputText("	Tipe Hp :"); !err {
				data.tipe = tipe
			}
			if keluhan, err := libs.InputText("	Keluhan :"); !err {
				data.keluhan = keluhan
			}
			if wa, err := libs.InputText("	Wa/No-HP :"); !err {
				data.wa = wa
			}
			if sandi, err := libs.InputText("	Kunci Layar :"); !err {
				data.sandi = sandi
			}
			if biaya, err := libs.InputNumber("	Biaya :"); !err {
				data.biaya = biaya
			}

			return
		}

		fmt.Println("Hp dengan ID ", id, "tidak ditemukan")
	}
}

func DeleteDiKembalilan() {

}
