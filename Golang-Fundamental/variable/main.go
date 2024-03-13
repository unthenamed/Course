package main

import "fmt"

func main() {
	var firstName string
	firstName = "abdul"
	var lastName string = "jalil"
	fmt.Println(firstName, lastName)
	fmt.Printf("hallo %s %s\n", lastName, firstName)

	var (
		age     int
		address string
	)
	age = 25
	address = "nusa"
	fmt.Printf("%s %s %d %s\n", firstName, lastName, age, address)

	var (
		bootcampName, bootcampNAddress = "enigma", "jakarta"
	)
	fmt.Println(bootcampName, bootcampNAddress)

	day, date, month := "senin", 12, "september"
	fmt.Println(day, date, month+" 2006")

	var number = 21
	number = 20
	const phi = 3.14
	fmt.Println(phi, number)

	firstName, lastName = lastName, firstName
	fmt.Println(firstName, lastName)
}
