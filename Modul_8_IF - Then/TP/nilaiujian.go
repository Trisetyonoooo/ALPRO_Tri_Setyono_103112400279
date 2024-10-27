package main

import "fmt"

func main () {
	var nilaiujian int
	// Inputan untuk memasukan nilai ujian
	fmt.Print("Masukan nilai ujian: ")
	fmt.Scan(&nilaiujian)
	// Memeriksa nilai ujian tersebut 70 atau lebih, jika iya program akan mencetak "Lulus"
	//  Jika tidak, program akan mencetak "Tidak Lulus"
	if nilaiujian >= 70 {
		fmt.Println("lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}