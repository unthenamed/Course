package library

import (
	"cowr/config"
	"fmt"
	"os"
)

func OpenFile() *os.File {
	file, err := os.OpenFile(config.PathFile, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return file
}
