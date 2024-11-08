package main

import "fmt"

type yarn struct {
	ne       int
	interval int
	mc       string
}

func main() {

	// penulisan pertama
	var rayon yarn
	rayon.ne = 30
	rayon.interval = 240
	rayon.mc = "base"

	fmt.Println(rayon)
	fmt.Println(rayon.ne)
	fmt.Println(rayon.interval)
	fmt.Println(rayon.mc)

	// penulisan kedua
	rayon40 := yarn{}
	rayon40.ne = 40
	rayon40.interval = 360
	rayon40.mc = "base"

	fmt.Println(rayon40)
	fmt.Println(rayon40.ne)
	fmt.Println(rayon40.interval)
	fmt.Println(rayon40.mc)

	//penulisan ketiga
	pe20 := yarn{20, 120, "jingway"}
	fmt.Println(pe20)

	//penulisan keempat
	sf20 := yarn{mc: "baseC", interval: 120, ne: 20}
	fmt.Println(sf20)

}
