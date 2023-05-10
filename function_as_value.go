package main

import "fmt"

func main() {
	bye := sayGoodbye
	result := bye("Mr")
	fmt.Println(result)
}

func sayGoodbye(nama string) string {
	return "Goodbye " + nama
}
