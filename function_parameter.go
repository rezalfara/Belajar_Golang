package main

import "fmt"

func main() {
	namaDepan := "Reza"
	sayHelloTo(namaDepan, "Alfara")
	sayHelloTo("Muhammad", "Reza")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Halo", firstName, lastName)
}
