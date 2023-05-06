package main

import "fmt"

func main() {
	nama := "Reza Alfara"

	switch nama {
	case "Reza":
		fmt.Println("Nama saya reza")
	case "Alfara":
		fmt.Println("Nama saya alfara")
	default:
		fmt.Println("Nama saya", nama)
	}

	//switch length := len(nama); length > 10 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah oke!")
	//}

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah OKE!")

	}

}
