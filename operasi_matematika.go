package main

import "fmt"

func main() {
	var result = 10 + 15
	fmt.Println(result)

	var a = 4
	var b = 5
	var c = a * b
	fmt.Println(c)

	//Augmented Assignment
	var i = 25
	i += 10 //	i = i + 10
	//fmt.Println(i)

	//Unary Operator
	i++ //	i = i + 1
	fmt.Println(i)

	var negatif = -10
	var positif = 10
	fmt.Println(negatif)
	fmt.Println(positif)

}
