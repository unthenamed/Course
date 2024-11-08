package main

import (
	"cowr/config"
	"cowr/library"
	"fmt"
)

func main() {
	config.Clear()
	library.CreateFile()
	file := library.OpenFile()
	defer file.Close()

	data := library.ReadFile(file)
	for i, v := range data {
		fmt.Println(i, v)
	}

	if text, err := config.InputText("Masukan User ID :"); !err {
		library.WriteFile(file, text)
		fmt.Println("done")
	}
}
