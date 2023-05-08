package main

import "fmt"

func main() {
	namaDepan, namaBelakang := getFullNama()
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}

func getFullNama() (firstName string, lastName string) {
	firstName = "Reza"
	lastName = "Alfara"

	return
}
