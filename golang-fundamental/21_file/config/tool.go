package config

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func recoverErr(e *bool) {
	if r := recover(); r != nil {
		fmt.Println("Recover :", r)
		*e = true
	}

}

func InputText(label string) (result string, err bool) {
	defer recoverErr(&err)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(label)
	scanner.Scan()
	result = scanner.Text()

	if result == "" {
		panic("Tidak boleh kosong")
	}

	return
}

func InputNumber(label string) (result int, err bool) {
	defer recoverErr(&err)

	fmt.Printf(label)
	fmt.Scanln(&result)

	if result == 0 {
		panic("Tidak boleh kosong")
	}

	return
}
