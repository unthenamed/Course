package main

import "fmt"

func inputUser() (result string) {
	fmt.Print("User ID :")
	fmt.Scanln(&result)
	return
}

// recover
func recoverPanic(h *bool) {
	if r := recover(); r != nil {
		fmt.Println("Perhatian :", r)
		*h = true
	}
}

func main() {
	nama := inputUser()

	if !apakahKosong(nama) {
		fmt.Println("User anda :", nama)
	}

	fmt.Println("done")
}

// panic
func apakahKosong(p1 string) (hasil bool) {
	defer recoverPanic(&hasil)
	if p1 == "" {
		panic("User tidak boleh kosong")
	}
	return
}
