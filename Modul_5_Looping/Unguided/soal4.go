package main

import "fmt"

func main () {
	var i, n, faktorial int
	// Meminta input dari pengguna untuk memasukan angka yang mau difaktorial
	fmt.Print("Masukan bilangan bulat non negatif: ")
	fmt.Scan(&n)
	// Menghitung Faktorial
	faktorial = 1
	for i = 1; i <= n; i ++ {
		faktorial = faktorial * i
	}
	// Menampilkan hasil faktorial
	fmt.Println("Faktorial dari", n, "adalah", faktorial)
}