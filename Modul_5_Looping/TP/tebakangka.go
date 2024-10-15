package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Deklarasikan variabel
	var angka int
	var kesempatan int
	var jawaban int
	const percobaan = 5
	// Untuk memilih angka acak antara 1 sampai 100
	rand.Seed(time.Now().UnixNano())
	jawaban = rand.Intn(100) + 1
	// Pengenalan program
	fmt.Println("Permainan tebak angka!")
	fmt.Println("Silahkan pilih angka dari 1 sampa 100.")
	fmt.Println("Anda diberikan 5 kesempatan untuk menebak angka tersebut")
	// Memberikan pengguna 5 kesempatan untuk menebak angka yang dimasukan
	for kesempatan = 1; kesempatan <= percobaan; kesempatan++ {
		fmt.Printf("Masukan angkamu: ")
		fmt.Scan(&angka)
		// Memeriksa tebakan pengguna apakah angkanya terlalu kecil / terlalu besar
		if angka < jawaban{
			fmt.Println("Angkanya Terlalu kecil!")
		} else if angka >  jawaban {
			fmt.Println("Angkanya Terlalu besar!")
		} else {
			fmt.Printf("Selamat! Anda berhasil menebak angka dalam %d percobaan!\n", kesempatan)
			return
		}
	}

	//  Menampilkan hasil jika pengguna gagal menebak setelah 5 kali percobaan
	fmt.Printf("Maaf, Anda telah kehabisan kesempatan. Angka yang benar adalah %d.\n", jawaban)
}
