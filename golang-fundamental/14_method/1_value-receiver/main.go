package main

import "fmt"

// struct
type ukuran struct {
	panjang int
	lebar   int
}

// method (fungsi yang terikat pada struct)
func (param ukuran) luas() {
	param.lebar += 10
	param.panjang += 10
	fmt.Println("method :", param)

}

func main() {
	tanah := ukuran{50, 20}
	tanah.luas()
	fmt.Println("main :", tanah)
}
