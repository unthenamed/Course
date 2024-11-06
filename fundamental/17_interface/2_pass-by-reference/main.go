package main

import "fmt"

type karyawan interface {
	getGaji() int
	updateGajiPokok(int)
	updateRek(int)
}

// Struct dan method untuk karyawan tetap
type tetap struct {
	nama      string
	gajiPokok int
	tunjangan int
	noRek     int
}

func (t tetap) getGaji() int {
	return t.gajiPokok + t.tunjangan
}
func (t *tetap) updateGajiPokok(g int) {
	t.gajiPokok = g
}
func (t *tetap) updateRek(r int) {
	t.noRek = r
}

// Struct dan method untuk karyawan kontrak
type kontrak struct {
	nama      string
	gajiPokok int
	noRek     int
}

func (k kontrak) getGaji() int {
	return k.gajiPokok
}
func (k *kontrak) updateGajiPokok(g int) {
	k.gajiPokok = g
}
func (k *kontrak) updateRek(r int) {
	k.noRek = r
}

// Struct dan method untuk karyawan harian
type harian struct {
	nama      string
	perHari   int
	hariKerja int
	noRek     int
}

func (h harian) getGaji() int {
	return h.perHari * h.hariKerja
}
func (h *harian) updateGajiPokok(g int) {
	h.perHari = g
}
func (h *harian) updateRek(r int) {
	h.noRek = r
}

func main() {
	// Upcasting
	var karyawan1 karyawan = &tetap{"jalil", 4000000, 200000, 59321728}
	var karyawan2 karyawan = &kontrak{"abdul", 300000, 2563719762}
	var karyawan3 karyawan = &harian{"khoironi", 135000, 26, 1637810002}

	karyawan1.updateGajiPokok(4500000)
	karyawan2.updateGajiPokok(3500000)
	karyawan3.updateGajiPokok(150000)

	fmt.Println(karyawan1.getGaji())
	fmt.Println(karyawan2.getGaji())
	fmt.Println(karyawan3.getGaji())

	fmt.Println(gajiSemuaKaryawan(karyawan1, karyawan2, karyawan3))
}

func gajiSemuaKaryawan(gaji ...karyawan) (total int) {
	for _, v := range gaji {
		total += v.getGaji()
	}
	return
}
