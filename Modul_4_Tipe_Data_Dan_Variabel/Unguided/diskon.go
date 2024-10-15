package main

import "fmt"

func main () {
	// deklarasi variabel
	var totalbawal, totalbakhir, diskon int					
	// inputan total belanja awal untuk user
	fmt.Print("Masukan total belanja awal :")				
	fmt.Scan(&totalbawal)
	// inputan diskon untuk user
	fmt.Print("Masukan diskon : ")							
	fmt.Scan(&diskon)
	// menghitung total belanja akhir
	totalbakhir= totalbawal - (totalbawal * diskon/100)		
	// Menampilkan hasil total belanja akhir
	fmt.Println("total belanja akhir adalah :", totalbakhir)
}