package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("masukan rating 1-10 : ")
	scanner.Scan()
	rating, _ := strconv.Atoi(scanner.Text())

	switch {
	case rating == 9:
		{
			fmt.Println("Unggul")
		}
	case rating == 8:
		{
			fmt.Println("Baik")

		}
		fallthrough // keyword untuk meneruskan ekselusi case selanjutnya
	case rating == 7:
		{
			fmt.Println("Cukup")
		}
	case rating == 6:
		{
			fmt.Println("Kurang")
		}
	default:
		{
			fmt.Println("Gagal")
		}
	}
}
