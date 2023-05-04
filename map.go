package main

import "fmt"

func main() {
	//Cara Pertama
	mahasiswa := map[string]string{
		"nama":    "Reza Alfara",
		"jurusan": "TIK",
		"prodi":   "Teknik Informatika",
	}

	mahasiswa["alamat"] = "Batuphat Timur"

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["nama"])
	fmt.Println(mahasiswa["alamat"], "\n")

	//Cara kedua
	//var hero map[string]string = make(map[string]string)
	hero := make(map[string]string)
	hero["nama"] = "Johson"
	hero["attack power"] = "50"
	hero["defence power"] = "95"

	fmt.Println(hero)
	fmt.Println(len(hero))

	delete(hero, "defence power")

	fmt.Println(hero)
	fmt.Println(len(hero))
}
