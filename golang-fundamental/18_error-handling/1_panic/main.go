package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 2
	result, err := pembagian(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func pembagian(par1, par2 int) (int, error) {
	if par2 == 0 {
		return 0, errors.New("pembagian dengan angka 0")
	}
	return par1 / par2, nil
}
