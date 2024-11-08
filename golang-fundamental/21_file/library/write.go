package library

import (
	"bufio"
	"os"
)

func WriteFile(file *os.File, input string) {

	write := bufio.NewWriter(file)
	write.WriteString(input + "\n")
	write.Flush()

}
