package main

import "fmt"

func main() {

	/* MAP */
	// untuk membuat map gunakan command
	// variable := map[type key][type value]

	userData := map[string]string{
		"username": "jalilk",
		"email":    "jalilkhoironi00@gmail.com",
	}
	fmt.Println(userData)

	// map juga bisa di buat dengan mengggunakan Make
	// untuk mengubah value dari sebuah map gunakan command
	// variableMap[keymap]="valueMap"
	var nilai = make(map[string]int)
	fmt.Println(nilai) //output map kosong

	nilai["indonesia"] = 80 //mengubah value
	nilai["inggris"] = 75
	nilai["mtk"] = 70
	fmt.Println(nilai) // map[indonesia:80 inggris:75 mtk:70]

	//untuk mengakses value dari map
	fmt.Println("Nilai indonesia :", nilai["indonesia"])

	//jika kita mengakses key tidak ada di map maka outputnya nilai default dari tiype value nya , ex int = 0
	fmt.Println("Nilai SKI :", nilai["ski"])

	// untuk menghapus value yang ada di dalam map
	delete(nilai, "mtk")
	fmt.Println(nilai)

	// untuk menginisialisasi map bisa menggunakan FORRANGE LOOPING
	for key, value := range nilai {
		fmt.Println("key :", key, "value :", value)

	}

	//jika ingin value saja bisa gunakan _
	for _, value := range nilai {
		fmt.Println("value :", value)

	}
}
