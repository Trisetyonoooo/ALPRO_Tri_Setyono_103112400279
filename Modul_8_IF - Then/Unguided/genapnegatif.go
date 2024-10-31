package main

import "fmt"

func main () {
	var bilangan int
	var teks string
	// Inputan untuk pengguna memasukan bilangannya
	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bilangan)
	teks = "bukan"
	// Memeriksa apakah bilangan tersebut bilangan negatif dan memeriksa apakah bilangannya bilangan genap
	// Jika bilangan negatif dan genap, hasilnya "genap negatif", jika tidak, hasilnya tetap "bukan".
	if bilangan < 0 && bilangan%2 == 0 {
		teks = "genap negatif"
	}
	fmt.Println(teks)
}