package main

import "fmt"

func main () {
	var i, a, b int
	var hasil int
	// Meminta input dari pengguna untuk memasukan 2 angka untuk dipangkatkan
	fmt.Print("Masukan dua bilangan bulat positif: ")
	fmt.Scan(&a, &b)
	hasil = 1
	// Melakukan pemangkatan menggunakan perkalian
	for i = 0; i < b; i++ {
		hasil = hasil * a
	}
	// Menampilkan hasil pemangkatan
	fmt.Println(hasil)
}