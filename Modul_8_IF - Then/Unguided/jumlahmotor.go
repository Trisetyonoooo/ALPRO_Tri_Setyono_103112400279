package main

import "fmt"

func main () {
	var orang int
	var motor int
	// Inputan untuk pengguna memasukan jumlah orang yang akan touring
	fmt.Print("Masukan jumlah orang yang akan touring: ")
	fmt.Scan(&orang)
	// Jika jumlah orang dapat dibagi dengan 2 maka jumlah motor adalah hasil bagi jumlah orang dengan 2
	// Tetapi ketika tidak bisa dibagi 2 maka jumlah motor dihitung adalah hasil bagi jumlah orang dengan 2 
	// dan menambahkan 1 untuk mengangkut orang yang tersisa
	if orang%2 == 0 {
		motor = orang / 2
	} else {
		motor = (orang / 2)+1
	}
	fmt.Println("Jumlah motor yang diperlukan adalah: ", motor)
}