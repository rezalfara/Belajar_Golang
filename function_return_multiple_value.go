package main

import "fmt"

func main() {
	_, lastName, title := getFullName()
	fmt.Println(lastName, title)
	fmt.Println(lastName)
	fmt.Println(title)
}

func getFullName() (string, string, string) {
	return "Reza", "Alfara,", "S.Tr.Kom"
}
