package main

import "fmt"

func main() {
	var hari = [...]string{
		"Minggu",
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
	}

	var slice1 = hari[3:5]   //membuat slice baru dari array
	fmt.Println(slice1)      //mencetak isi slice
	fmt.Println(len(slice1)) //menghitung panjang slice
	fmt.Println(cap(slice1)) //menghitung kapasitas slice

	//hari[3] = "Wednesday"
	//fmt.Println(slice1)

	//slice1[0] = "RabuLagi"
	//fmt.Println(hari, "\n")

	//append slice
	var slice2 = hari[2:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Natal")
	fmt.Println(slice3)
	slice3[3] = "Berkah"
	fmt.Println(slice3)

	//membuat slice baru dengan fungsi make
	sliceBaru := make([]string, 2, 5)
	sliceBaru[0] = "Reza"
	sliceBaru[1] = "Alfara"

	fmt.Println("\n", sliceBaru)
	fmt.Println(len(sliceBaru))
	fmt.Println(cap(sliceBaru))

	//menyalin slice
	salinSlice := make([]string, len(sliceBaru), cap(sliceBaru))
	copy(salinSlice, sliceBaru)
	fmt.Println(salinSlice)

}
