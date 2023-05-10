package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Reza", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHelloWithFilter(nama string, filter Filter) {
	nameFiltered := filter(nama)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(nama string) string {
	if nama == "Anjing" {
		return "..."
	} else {
		return nama
	}
}
