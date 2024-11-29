package main

import "fmt"

func main() {
	// Variabel
	var username, password string
	// Variabel untuk menyimpan password dan username yang benar
	correctUsername := "Admin"
	correctPassword := "Admin"
	// Variabel untuk menyimpan jumlah percobaan yang gagal
	jumlahPercobaan := 0
	// Perulangan untuk memberi kesempatan pada user memasukan username dan password
	for {
		// Meminta inputan username dan password
		fmt.Print("Masukan username dan password: ")
		fmt.Scan(&username, &password)
		// Mememriksa apakah username dan password yang dimasukan benar
		if username == correctUsername && password == correctPassword {
			// Jika benar program akan keluar dari perulangan dan menghitung jumlah percobaan yang gagal
			fmt.Printf("%d percobaan gagal login.\n", jumlahPercobaan)
			break
		} else {
			// Jika salah, user akan diberikan kesempatan lagi sampai memasukan username dan password yang benar
			jumlahPercobaan++
		}
	}
}
