package main

import "fmt"

func main () {
	var totalbawal, totalbakhir, diskon int							// deklarasi variabel

	fmt.Print("Masukan total belanja awal :")						// inputan total belanja awal untuk user
	fmt.Scan(&totalbawal)
	fmt.Print("Masukan diskon : ")									// inputan diskon untuk user
	fmt.Scan(&diskon)

	totalbakhir= totalbawal - (totalbawal * diskon/100)				// menghitung total belanja akhir

	fmt.Println("total belanja akhir adalah :", totalbakhir)
}