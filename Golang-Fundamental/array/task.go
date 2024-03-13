package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input = bufio.NewScanner(os.Stdin)
	var kapasitas int
	input.Scan()
	kapasitas, _ = strconv.Atoi(input.Text())

	var hasil = make([]int, kapasitas)

	input.Scan()
	elements := input.Text()
	elementList := strings.Split(elements, " ")

	for _, str := range elementList {
		elements, _ := strconv.Atoi(str)
		hasil = append(hasil, elements)
	}

	for _, numbers := range hasil {
		if numbers%2 == 0 {
			if numbers != 0 {
				fmt.Println(numbers)
			}
		}
	}
}
