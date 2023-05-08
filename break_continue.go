package main

import "fmt"

func main() {

	//Break
	for i := 1; i <= 10; i++ {
		if i == 7 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}

	//Continue
	for j := 0; j <= 10; j++ {
		if j%2 != 0 {
			continue
		}
		fmt.Println("Looping ke", j)
	}
}
