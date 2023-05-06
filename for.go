package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Looping ke", counter)
		counter++
	}

	//Looping dengan statement
	//for counter := 1; counter <= 10; counter++ {
	//	fmt.Println("Looping ke", counter)
	//}
	fmt.Println()

	slice := []string{"saya", "kamu", "dia", "mereka", "kita"}

	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}

	for _, value := range slice {
		//fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	fmt.Println()

	person := make(map[string]string)
	person["nama"] = "Reza Alfara"
	person["title"] = "Programmer"

	for kunci, nilai := range person {
		fmt.Println(kunci, "=", nilai)
	}

}
