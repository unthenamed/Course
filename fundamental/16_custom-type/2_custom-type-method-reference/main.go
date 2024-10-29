package main

import "fmt"

type Celsius float64

func (param *Celsius) keFarenheit() {
	*param = *param*9/5 + 32
}

func main() {

	var suhu Celsius = 40.5
	suhu.keFarenheit()
	fmt.Println(suhu)
}
