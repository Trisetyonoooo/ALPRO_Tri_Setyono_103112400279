package main

import "fmt"

func main() {
	// Deklarasikan variabel n dan hasil
	var n, hasil int
	// Inputan untuk memasukan nilai n
	fmt.Print("Masukan nilai n: ")
	fmt.Scan(&n)
	// Menjumlahkan semua angka dari 1 hingga n dengan rumus for looping
	hasil = 0
	for i := 1; i <= n; i++{
		hasil += i
	}
	fmt.Println("Jumlah dari deret angka 1 hingga", n ,"adalah", hasil)

}