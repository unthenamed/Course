package main

import "fmt"

func main() {
	var positiveNumber uint8 = 90
	var negativeNumber int = -90
	var decimalNumber = 3.15
	var exist = true
	var message = "halo"
	var otheMessage = `saya 
	dsc
	c`
	fmt.Printf(" %d\n%d\n%.2f\n%t\n%s\n%s\n", positiveNumber, negativeNumber, decimalNumber, exist, message, otheMessage)
	fmt.Println(len(otheMessage))
}
