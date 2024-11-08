package main

import "fmt"

func main() {

	var name string = "jalil"
	var myName *string = &name

	*myName = "abdul"

	fmt.Println(name)
	fmt.Println(*myName)
}
