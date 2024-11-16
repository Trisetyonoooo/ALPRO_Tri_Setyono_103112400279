package main

import "fmt"

func main() {

	// Variabel
	var b, prima int
	var hasil bool

	// Input bilangan
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&b)

	// Validasi input bahwa nilai b harus lebih besar dari 1
	if b > 1 { 
	// Menampilkan faktor bilangan
		fmt.Print("Faktor: ")
	// Memeriksa setiap angka dari 1 hingga b untuk melihat apakah angka tersebut adalah faktor dari b
		for i := 1; i <= b; i++ { 
	// Mengecek apakah i adalah faktor dari b. Jika hasilnya 0, maka i adalah faktor dari b
			if b%i == 0 {
				fmt.Print(i, " ")
	// Menghitung jumlah faktor dari b
				prima++
			}
		}

	// Menentukan apakah bilangan prima
		hasil = prima == 2
		fmt.Println() 
		fmt.Println("Prima:", hasil)
	} else {
		// Pesan kesalahan untuk bilangan <= 1
		fmt.Println("Masukkan bilangan bulat lebih dari 1!")
	}
}
