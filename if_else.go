package main

import "fmt"

func main() {

	var angka = 35

	if angka == 10 {
		fmt.Println("Angka adalah 10")
	} else if angka == 0 {
		fmt.Println("Angka adalah 0")
	} else {
		fmt.Println("Angka adalah ", angka)
	}

	//Short statement if
	var nama = ""
	//length := len(nama)

	if length := len(nama); length < 1 {
		fmt.Println("Nama terlalu pendek")
	} else {
		fmt.Println("Nama : ", nama)
	}

}
