package library

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var Location, _ = time.LoadLocation("Asia/Jakarta")

func recoverPanic(result *bool) {
	if r := recover(); r != nil {
		fmt.Println(r)
		*result = true
	}
}

func IsNil(err error) (result bool) {
	defer recoverPanic(&result)
	if err != nil {
		panic(err)
	}
	return
}
func IsNilCustom(err error, label string) (result bool) {
	defer recoverPanic(&result)
	if err != nil {
		panic(errors.New(label))
	}
	return
}

func InputText(label string) (result string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(label)
	scanner.Scan()
	result = scanner.Text()
	return
}

func InputNumber(label string) (result int) {
	fmt.Printf(label)
	fmt.Scanln(&result)
	return
}

func NewCustomer() (result Customer) {
	result.Customer_id = InputNumber("Input ID :")
	result.Name = InputText("Input Name :")
	result.Phone = InputText("Input phone :")
	result.Address = InputText("Input Address :")
	return

}
func NewService() (result Service) {
	result.Id = InputNumber("Input ID :")
	result.Name = InputText("Input Service Name :")
	result.Unit = InputText("Input Unit :")
	result.Price = InputNumber("Input price :")
	return

}
