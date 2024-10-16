package main

import (
	"fmt"
)

func main() {
	var i, n int
	var jumlah int
	
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)
	// Menghitung jumlah dari 1 sampai n
	jumlah = 0
	for i = 1; i <= n; i++ {
		jumlah += i
	}
	// Menampilkan hasil penjumlahan
	fmt.Println(jumlah)
}