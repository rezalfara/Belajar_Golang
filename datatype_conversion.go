package main

import "fmt"

func main() {
	var nilai32 int32 = 130
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8, "\n")

	var nama = "Reza Alfara"
	var z = nama[2]
	fmt.Println(z)
	var zString = string(z)
	fmt.Println(zString)
}
