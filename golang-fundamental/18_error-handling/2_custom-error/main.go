package main

import "fmt"

type siswa struct {
	nama   string
	alamat string
	id     int
}

type idNotFund struct {
	id int
}

func (i *idNotFund) Error() string {
	return fmt.Sprintf("siswa dengan id %d tidak di temukan", i.id)
}

func main() {
	siswa1 := siswa{"abdul", "nusa", 1}
	siswa2 := siswa{"jalil", "nusamasjid", 2}
	siswa3 := siswa{"khoironi", "nusatengah", 3}
	dataSiswa := []siswa{siswa1, siswa2, siswa3}

	fmt.Println(cariSiswa(dataSiswa, 3))

}

func cariSiswa(data []siswa, id int) (*siswa, error) {
	for _, value := range data {
		if value.id == id {
			return &value, nil
		}
	}
	return nil, &idNotFund{id: id}
}
