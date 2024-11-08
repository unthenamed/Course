package main

func main() {
	array := [5]int{
		1,
		1,
		1,
		1,
		1}

	slice1 := array[:2]
	slice2 := array[2:]

	apSlice1 := append(slice1, 2)
	apSlice2 := append(slice2, 3)

}
