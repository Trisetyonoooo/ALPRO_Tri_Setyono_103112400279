package main

import "fmt"

func main() {
	var totalbelanja float32
	var asesmen bool
	var totalbelanjaakhir, diskon float32

	fmt.Print("Masukan total belanja awal: ")
	fmt.Scan(&totalbelanja)
	fmt.Print("Apakah sedang mengikuti asesmen CLO?(true/false): ")
	fmt.Scan(&asesmen)

	if asesmen {
		diskon = 0.35
	} else {
		diskon = 0.0
	}
	totalbelanjaakhir = totalbelanja * (1 - diskon)
	fmt.Printf("Total belanja setelah diskon: %.0f\n", totalbelanjaakhir)
}
