package main 

import "fmt"

func main() {
	// Deklarasi variabel
	var angka int
	// Perulangan untuk menebak angka dari 1 sampai 10
	for {
		// Meminta Input dari pengguna
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&angka)
		// Jika angka yang dimasukan 7 maka program akan keluar dari perulangan
		if angka == 7 {
			fmt.Println("Selamat, tebakan Anda benar!.")
			break
		// jika salah, program akan melanjutkan perulangan sampai memasukan tebakan yang benar	
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")		
		}
	}
}