package main

import "fmt"

func main() {
	var harga, totalbelanja int
	// Perulangan untuk memasukan harga barang dan akan menghitung total belanja ketika user memasukan harga 0
	for {
		// Meminta input harga barang
		fmt.Print("Masukan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&harga)
		// Menghitung total belanja dari harga barang
		totalbelanja += harga
		// jika user memasukan harga barang 0 maka program akan keluar dari perulangan
		if harga == 0 {
			break
		}
	}
	// Menampilkan total belanja setelah perulangan selesai
	fmt.Println("Total belanja Anda: ", totalbelanja)
}