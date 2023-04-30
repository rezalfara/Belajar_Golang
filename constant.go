package main

import "fmt"

func main() {
	//Variabel bisa diubah nilainya
	const nama string = "Reza Alfara"
	fmt.Println(nama)

	var umur = 22
	fmt.Println(umur)

	umur = 23
	fmt.Println(umur)

	//Constant nilainya tetap (tidak bisa diubah)
	const agama = "Islam"
	fmt.Println(agama)

	//agama = "agama_lain"

	//multi constant
	const (
		NIM     = 1857301003
		jurusan = "TIK"
		prodi   = "Teknik Informatika"
	)
	fmt.Println(NIM)
	fmt.Println(jurusan)
	fmt.Println(prodi)
}
