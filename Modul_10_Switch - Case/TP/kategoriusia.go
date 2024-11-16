package main 

import "fmt"

func main () {
	// Variabel
	var usia int
	// Meminta input dari pengguna
	fmt.Print("Masukan usia anda: ")
	fmt.Scan(&usia)
	// Switch Case untuk menentukan kategori usia berdasarkan input usia dari pengguna
	switch {
	case usia >= 0 && usia <= 12 :
			fmt.Println("Kategori: Anak-anak")
	case usia >= 13 && usia <= 17 :
			fmt.Println("Kategori: Remaja")
	case usia >= 18 && usia <= 64 :
			fmt.Println("Kategori: Dewasa")
	default :
			fmt.Println("Kategori: Lansia")
	}
}