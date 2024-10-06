package main

import "fmt"

func main() {
	var r float64

	fmt.Print("Masukan jari jari Lingkaran =") // Input Jari jari
	fmt.Scan(&r)

	luas := 3.14 * r * r // Rumus menghitung Luas Lingkaran
	keliling := 2 * 3.14 * r // Rumus menghitung Keliling Lingkaran

	// Menampilkan hasil 
	fmt.Println("Luas Lingkaranya adalah =", luas,"cm²")
	fmt.Println("Keliling lingkaranya adalah =", keliling,"cm")
}