package main

import "fmt"

func main () {
	// Deklarasi Variabel
	var hargabarang, jumlahbarang int
	var totalbelanja float64
	var tambahbarang, namabarang string

	fmt.Println("<==== Aplikasi Kasir Sederhana ====>")
	for {
		// Input nama barang
		fmt.Print("Masukan nama barang: ")
		fmt.Scan(&namabarang)
		// Input jumlah barang
		fmt.Print("Masukan jumlah barang: ")
		fmt.Scan(&jumlahbarang)
		// Input harga barang
		fmt.Print("Masukan harga barang: Rp.")
		fmt.Scan(&hargabarang)
		// Menghitung total belanja
		totalbelanja += float64(hargabarang)* float64(jumlahbarang)
		// Input untuk menanyakan apakah ada tambahan barang lagi
		fmt.Print("Apakah ada tambahan barang lagi? (ya/tidak): ")
		fmt.Scan(&tambahbarang)
		// Jika tidak ada tambahan barang, maka program akan keluar dari perulangan
		if tambahbarang == "tidak"{
			break
		}
	}
	// Menampilkan total akhir belanja
	fmt.Printf("\nTotal belanja anda: Rp.%.2f\n", totalbelanja)
	fmt.Println("<=== Terima kasih sudah berbelanja!!! ===>")
}