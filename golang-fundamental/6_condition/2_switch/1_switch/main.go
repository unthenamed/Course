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
	fmt.Print("Masukan rating 1-10 :")
	scanner.Scan()
	rating, _ := strconv.Atoi(scanner.Text())

	switch rating {
	case 9:
		{
			result = "Mantap"
		}
	case 8, 7:
		{
			result = "Baik"
		}
	default:
		{
			result = "Kurang"
		}
	}

	fmt.Println(result)
}
