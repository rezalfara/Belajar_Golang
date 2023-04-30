package main

import "fmt"

func main() {
	//Cara Pertama
	var nama_depan string
	nama_depan = "Reza"
	fmt.Println(nama_depan)

	//Cara Kedua
	var nama_belakang string = "Alfara"
	fmt.Println(nama_belakang)

	//Cara Ketiga
	var alamat = "Batuphat Timur"
	fmt.Println(alamat)

	//Cara Keempat
	umur := 22
	fmt.Println(umur)

	//Multiple Variable
	var (
		firstname = "nama depan"
		lastname  = "nama belakang"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}
