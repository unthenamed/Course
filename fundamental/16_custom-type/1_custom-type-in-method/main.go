package main

import "fmt"

type Celsius float64

func (param Celsius) toFarenheit() float64 {
	return float64(param*9/5 + 32)
}

func main() {
	var temperatur Celsius = 32.4

	fmt.Println(temperatur.toFarenheit())
}
