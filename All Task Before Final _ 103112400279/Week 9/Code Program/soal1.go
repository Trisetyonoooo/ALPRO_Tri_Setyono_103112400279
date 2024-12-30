package main

import "fmt"

func main() {
	var N, mobil int

	// Input jumlah mahasiswa
	fmt.Scan(&N)

	// Menentukan jumlah mobil yang diperlukan dengan menggunakan else if
	if N <= 7 {
		mobil = 1
	} else if N <= 14 {
		mobil = 2
	} else if N <= 21 {
		mobil = 3
	} else {
		mobil = (N + 6) / 7
	}

	// Output jumlah mobil yang diperlukan
	fmt.Println(mobil)
}
