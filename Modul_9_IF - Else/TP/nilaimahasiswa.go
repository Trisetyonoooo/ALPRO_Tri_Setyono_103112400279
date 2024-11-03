package main

import "fmt"

func main() {
	var nilai int
	// Inputan untuk pengguna memasukan nilainya
	fmt.Print("Masukan nilai anda: ")
	fmt.Scan(&nilai)
	// Memeriksa nilai yang diinputkan dan menentukan indeks berdasarkan rentang nilai yang sesuai menggunakan percabangan
	if nilai >90 {
		fmt.Println("Indeks A")
	} else if nilai >= 80 && nilai <= 90 {
		fmt.Println("Indeks AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Indeks B")
	} else if nilai < 70 {
		fmt.Println("Indeks C")
	}
}