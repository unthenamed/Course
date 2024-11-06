package library

import (
	"bufio"
	"os"
)

func ReadFile(file *os.File) (result []string) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}
