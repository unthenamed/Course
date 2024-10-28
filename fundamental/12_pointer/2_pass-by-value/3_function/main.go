package main

import "fmt"

func main() {

	pointer := 25

	increase(pointer)
	fmt.Println(pointer, "alamat memori :", &pointer, "di main")

}

func increase(param int) {
	param++
	fmt.Println(param, "alamat memori :", &param, "di function")
}
