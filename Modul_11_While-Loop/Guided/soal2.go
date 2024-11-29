package main

import "fmt"	

func main() {
	var token string

	fmt.Print("Masukan password: ")
	fmt.Scan(&token)

	for token != "12345abcde" {
		fmt.Println("Password salah, coba lagi.")
		fmt.Print("Masukan password: ")
		fmt.Scan(&token)
	}

	fmt.Println("Selamat, Anda berhasil login.")
}