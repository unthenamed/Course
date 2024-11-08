package library

import (
	"cowr/config"
	"fmt"
	"os"
)

func CreateFile() {
	_, errorPath := os.Stat(config.PathFile)

	if os.IsNotExist(errorPath) {
		fmt.Println("File tidak ditemukan ", config.PathFile)

		file, errorCreate := os.Create(config.PathFile)
		if errorCreate != nil {
			fmt.Println("Error :", errorCreate)
			return
		}

		defer file.Close()
		fmt.Println("File Berhasil dibuat ")
		return
	}

	fmt.Println("File sudah ada ")
}
