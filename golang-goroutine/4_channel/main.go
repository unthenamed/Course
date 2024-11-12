package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// Deklarasi channel bertipedata string
	cString := make(chan string)

	wg.Add(2)
	go getNotif(cString)
	go sendNotif(cString, "Selamat pagi")
	wg.Wait()
}

func sendNotif(channel chan string, notif string) {
	//pengirim
	channel <- notif
	wg.Done()
}
func getNotif(channel chan string) {
	// penrima menunggu hingga ada pengirim
	notif := <-channel
	fmt.Println(notif)
	wg.Done()
}
