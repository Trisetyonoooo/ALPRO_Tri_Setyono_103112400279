package main

import "fmt"

func main() {
	var donasi, target, totalDonasi, jumlah int
	// Meminta user memasukkan target donasi
	fmt.Print("Masukan target donasi: ")
	fmt.Scan(&target)
	// Perulangan untuk memasukan donasi sampai target donasi tercapai
	for selesai := false; !selesai; {
		// Meminta user memasukkan donasi
		fmt.Print("Masukan donasi: ")
		fmt.Scan(&donasi)
		totalDonasi += donasi
		selesai = totalDonasi >= target
		jumlah++
		fmt.Printf("Total donasi: %d\n", totalDonasi)
	}
	// Outpus setelah target donasi tercapai
	fmt.Printf("Target donasi tercapai!\nTotal donasi: %d\nJumlah donatur: %d\n", totalDonasi, jumlah)
}