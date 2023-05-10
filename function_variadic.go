package main

import "fmt"

func main() {
	result := sumAll(2, 4, 3)
	fmt.Println(result)
	fmt.Println(sumAll(10, 12, 10))
	fmt.Println(sumAll())

	slice := []int{23, 44, 10}
	hasil := sumAll(slice...)
	fmt.Println(hasil)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
