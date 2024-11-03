package main

import "fmt"

func main() {
	var umur int
	var kewarganegaraan string
	// Inputan untuk pengguna memasukan umur dan kewarganegaraanya
	fmt.Print("Masukan umur anda: ")
	fmt.Scan(&umur)
	fmt.Print("Masukan kewarganegaraan anda (WNI/WNA): ")
	fmt.Scan(&kewarganegaraan)
	// Memeriksa dua syarat utama yaitu usia dan kewarganegaraan yang diinputkan
	// Jika pengguna berusia 17 tahun atau lebih dan merupakan WNI, maka mereka dapat mengikuti pemilu
	// Jika salah satu dari syarat tersebut tidak terpenuhi, program akan menampilkan alasan yang membuat user tidak bisa mengikuti pemilu
	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena belum cukup umur")
		}
		if kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena bukan WNI")
		}
	}
}