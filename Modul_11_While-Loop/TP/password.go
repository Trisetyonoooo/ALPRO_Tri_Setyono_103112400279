package main

import "fmt"

func main () {
	var password string
	// Variabel untuk menyimpan password yang benar
	var passwordbenar = "Agussedihbanget"
	// Perulangan untuk memberi kesempatan pada user memasukan password sebanyak 3 kali
	for i := 1; i <= 3; i++ {
		// Meminta inputan password dari user
		fmt.Print("Masukan password terlebih dahulu: ")
		fmt.Scan(&password)
		// Mememriksa apakah password yang dimasukan benar atau tidak
		if password == passwordbenar {
			fmt.Println("Anda berhasil Login!!!")
			// jika benar program akan keluar dari perulangan 
			break
		} else {
			// Ketika user salah memasukan password, user akan diberikan kesempatan lagi
			fmt.Println("Password yang anda masukan salah, coba lagi!!!")
		}
		// Jika user gagal memasukan password yang benar sampai percobaan ke 3 maka program akan berhenti
		if i == 3{
			fmt.Println("Login ditolak")
		}
	}
}