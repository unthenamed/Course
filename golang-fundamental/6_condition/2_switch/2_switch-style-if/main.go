package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan rating 1-10 : ")
	scanner.Scan()
	rating, _ := strconv.Atoi(scanner.Text())

	switch {
	case rating == 10:
		{
			result = "Terbaik"
		}
	case rating >= 9 && rating < 10:
		{
			result = "Baik"
		}
	default:
		{
			result = "Cukup"
		}
	}

	fmt.Println(result)
}
