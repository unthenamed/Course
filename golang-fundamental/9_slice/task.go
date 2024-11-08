package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func evenNames(slice []string) []string {
	var evenNames []string
	for _, name := range slice {
		if len(name)%2 == 0 {
			evenNames = append(evenNames, name)
		}
	}
	return evenNames
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	slice := strings.Split(x, " ")
	names := evenNames(slice)
	for _, v := range names {
		fmt.Printf("%s ", v)
	}

}
