package main

import "fmt"

func main() {
	var makanan = [4]string{"kue", "bakwan", "es", "susu"}
	for i := 0; i < len(makanan); i++ {
		fmt.Println(makanan[i])
	}
}
