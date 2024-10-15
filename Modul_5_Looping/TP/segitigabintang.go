package main

import "fmt"

func main() {
	var jumlahbaris int
	// Inputan untuk memasukan jumlah baris
	fmt.Print("Masukkan jumlah baris segitiga bintang: ")
	fmt.Scan(&jumlahbaris)
	// Menentukan jumlah baris segitiga bintang yang akan dicetak menggunakan for looping
	for i := 1; i <= jumlahbaris; i++ {
		for baris := 1; baris <= i; baris++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
