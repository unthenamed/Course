package main

import "fmt"

type Celsius float64

type Pasien struct {
	nama string
	usia int
	Celsius
}

func main() {
	pasien1 := Pasien{"abdul", 25, 34.6}

	fmt.Printf("%+v", pasien1)
}
