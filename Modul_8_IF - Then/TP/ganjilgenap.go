package main

import "fmt"

func main () {
	var angka int
	// Inputan untuk memasukan angka
	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)
	// Memeriksa apakah angka habis dibagi 2 menggunakan operator modulus % 
	// Jika sisa bagi sama dengan 0, maka angka tersebut genap. jika tidak maka angka tersebut ganjil
	if angka%2 == 0 {
		fmt.Println("Angka adalah genap")
	} else {
		fmt.Println("Angka adalah ganjil")
	}
}