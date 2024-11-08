package main

import "fmt"

func main() {

	user := map[string]string{
		"name":  "abdul jalil",
		"email": "jalil@mail.com",
	}

	user["alamat"] = "kp.Nusa"

	fmt.Println(user)
}
