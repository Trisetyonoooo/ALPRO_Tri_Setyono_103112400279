package main

import "fmt"

func main() {
	var bilangan int
	var prima bool
	// Meminta input untuk memasukan bilangan dari user
	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bilangan)
	// Perulangan untuk memeriksa setiap bilangan dari 2 hingga (bilangan - 1)
	for i := 2; i <= bilangan; i++ {
		prima = true
		// Perulangan untuk memeriksa faktor bilangan dari i
		for j := 2; j*j <= i; j++ {
			// Jika i habis dibagi dengan j maka i bukan bilangan prima
			if i%j == 0 {
				prima = false
				break
			}
		}
		// Menampilkan output bilangan prima, jika i adalah bilangan prima
		if prima {
			fmt.Print(i, " ")
		}
	}	
}
