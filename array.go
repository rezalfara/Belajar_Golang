package main

import "fmt"

func main() {
	//membuat array dengan cara manual
	var nama [2]string
	nama[0] = "Reza"
	nama[1] = "Alfara"

	fmt.Println(nama[0])
	fmt.Println(nama[1])

	//membuat array secara langsung
	var values = [3]int{20, 40, 60}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//mengetahui panjang array
	fmt.Println(len(nama))
	fmt.Println(len(values))

}
