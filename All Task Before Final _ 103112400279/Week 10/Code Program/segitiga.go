package main

import "fmt"

func main() {
	var sisi1, sisi2, sisi3 int

	// Input tiga sisi segitiga
	fmt.Scan(&sisi1)
	fmt.Scan(&sisi2)
	fmt.Scan(&sisi3)

	// Menentukan jenis segitiga
	if sisi1 == sisi2 && sisi2 == sisi3 {
		// Semua sisi sama
		fmt.Println("Segitiga Sama Sisi")
	} else if sisi1 == sisi2 || sisi2 == sisi3 || sisi1 == sisi3 {
		// Dua sisi sama
		fmt.Println("Segitiga Sama Kaki")
	} else {
		// Semua sisi berbeda
		fmt.Println("Segitiga Sembarang")
	}
}
