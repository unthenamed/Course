package main

import "fmt"

func main() {
	defer fmt.Println("SATU")
	defer fmt.Println("DUA")
	defer fmt.Println("TIGA")
	defer fmt.Println("EMPAT")
}
