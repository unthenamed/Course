package main

import "fmt"

func main() {
	var myArr = [4]string{"satu", "dua", "tiga", "empat"}
	fmt.Println(myArr)

	myArr[2] = "duadua"
	fmt.Println(myArr)
}
