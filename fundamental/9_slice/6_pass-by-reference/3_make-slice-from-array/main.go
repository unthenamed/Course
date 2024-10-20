package main

import "fmt"

func main() {
	// Array
	fmt.Println("Array")
	nameArray := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu"}
	fmt.Println(nameArray)
	fmt.Println()

	//	[senin selasa rabu kamis jumat sabtu]

	///////////////////////////////////////////////
	// Make slice from array
	nameSlice := nameArray[:3]
	fmt.Println(nameSlice)
	fmt.Printf("len %d cap  %d\n", len(nameSlice), cap(nameSlice))
	fmt.Println()
	//	[senin selasa]
	//	len 2 cap  6

	// mencoba perubahan data pada slice dan array
	nameSlice[0] = "Slice"
	nameArray[2] = "Array"

	fmt.Println(nameArray)
	fmt.Println(nameSlice)

	///////////////////////////////////////////////
	fmt.Println()
	nameSlice2 := nameArray[4:]
	fmt.Println(nameSlice2)
	nameSlice2 = append(nameSlice2, "dadu")
	nameSlice2[1] = "roni"
	fmt.Println(nameSlice2)
	fmt.Println(nameArray)
	fmt.Printf("len %d cap  %d\n", len(nameSlice2), cap(nameSlice2))

}
