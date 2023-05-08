package main

import "fmt"

func main() {
	hasil := getHello("Reza")
	fmt.Println(hasil)

	fmt.Println(getHello(""))
}

func getHello(nama string) string {
	if nama == "" {
		return "Halo Bro!"
	} else {
		return "Halo " + nama
	}
}
