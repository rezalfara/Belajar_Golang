package main

import "fmt"

func main() {
	//var ujian = 75
	//var absensi = 75

	var lulusUjian = true
	var lulusAbsensi = true
	fmt.Println("Lulus Ujian\t= ", lulusUjian)
	fmt.Println("Lulus Absensi\t= ", lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println("Lulus\t\t= ", lulus)
}
