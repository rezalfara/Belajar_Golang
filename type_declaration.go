package main

import "fmt"

func main() {
	type NIM string

	var Reza NIM = "1857301003"
	fmt.Println(Reza)
	fmt.Println(NIM("12341234"))
}
