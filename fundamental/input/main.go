package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var input = bufio.NewScanner(os.Stdin)

func main() {
	var firstName string
	var fullName string
	var age int
	var code int

	fmt.Scanln(&firstName)
	fmt.Scanln(&code)
	fmt.Printf("%s#%d", firstName, code)

	input.Scan()
	fullName = input.Text()
	fmt.Println(fullName)

	input.Scan()
	age, _ = strconv.Atoi(input.Text())
	fmt.Println(age)

}
