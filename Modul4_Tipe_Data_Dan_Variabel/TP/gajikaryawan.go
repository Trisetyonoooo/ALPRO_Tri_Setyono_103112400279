package main

import "fmt"

func main() {
	var (
		seminggu float64
		jam float64
		total float64
	)

	fmt.Print("Masukan jumlah jam kerja dalam seminggu : ") // Input jumlah jam kerja dalam seminggu
	fmt.Scan(&seminggu)
	fmt.Print("Masukan Upah per jam : ") // Input Upah per jam
	fmt.Scan(&jam)

	if seminggu > 40 {
		total = (40 * jam) + (seminggu - 40) * 1.5 * jam //Rumus menhitung gaji jika jam kerja lebih dari 40 jam
	} else {
		total = seminggu * jam // Rumus menghitung gaji jika jam kerja akurang dari atau sama dengan 40 jam
	}

	fmt.Println("Total Gaji Bulanan : Rp", total) // Menampilkan total gaji bulanan
}	