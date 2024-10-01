package main

import "fmt"

func main() {
	var (
		nama  string
		NIM   int
		kelas string
	)

	println("Biodata Mahasiswa")

	print("Masukkan Nama =")
	fmt.Scanln(&nama)

	print("Masukkan NIM =")
	fmt.Scanln(&NIM)

	print("Masukkan Kelas =")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan nama saya adalah", nama , "Dari kelas",kelas , "Dengan NIM", NIM)
}