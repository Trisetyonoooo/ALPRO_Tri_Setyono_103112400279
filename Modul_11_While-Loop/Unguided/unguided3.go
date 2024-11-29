package main

import "fmt"		

func main() {
	// Variabel
	var x, y int
	// Inisialisasi variabel 
	hasil := 0
	// Input untuk memasukan bilangan x dan bilangan y
	fmt.Print("Masukan nilai x dan y: ")
	fmt.Scan(&x, &y)
	// Memastikan bahwa x harus lebih besar dari sama dengan y
	if x < y {
		fmt.Println("Bilangan x harus lebih besar dari bilangan y.")
		return
	}
	// Perulangan untuk menghitung integer division
	for x >= y {
		// Mengurangi x dengan y sampai x kurang dari y
		x -= y
		// Setiap kali mengurangi x dengan y maka hasil akan bertambah
		hasil++
	}
	// Menampilkan hasil integer division
	fmt.Println(hasil)
}