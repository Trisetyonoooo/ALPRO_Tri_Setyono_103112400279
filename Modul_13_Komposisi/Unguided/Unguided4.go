package main

import "fmt"

func main() {
	// Deklarasi Variabel
	var namabunga string
	var pita string
	var jumlahBunga int
	// Inisialisasi Variabel
	jumlahBunga = 0
	// Perulangan tak terbatas, yang akan terus meminta input nama bunga hingga pengguna memasukkan kata "SELESAI".
	for {
		fmt.Print("Bunga", jumlahBunga+1, ": ")
		fmt.Scan(&namabunga)

		if namabunga== "SELESAI" {
			break
		}

		pita = pita + namabunga + " - "
		jumlahBunga++
	}
	// Output nama bunga yang diinputkan
	fmt.Println("Pita: ", pita)
	// Output jumlah bunga yang diinputkan
	fmt.Println("Bunga: ", jumlahBunga)
}
