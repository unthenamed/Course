package main

import (
	"fmt"
	"sync"
)

type Akun struct {
	id     int
	kontak string
	saldo  int
	mutex  sync.Mutex
	wg     sync.WaitGroup
}

func (a *Akun) topup(n int) {
	a.mutex.Lock()
	a.saldo += n
	a.mutex.Unlock()
	a.wg.Done()
}

func main() {

	pelanggan1 := Akun{id: 1, kontak: "082310607461", saldo: 10}

	for i := 0; i < 100; i++ {
		pelanggan1.wg.Add(1)
		go pelanggan1.topup(50)
	}
	pelanggan1.wg.Wait()
	fmt.Println(pelanggan1.saldo)

}
