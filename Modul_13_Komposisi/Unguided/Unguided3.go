package main 

import "fmt"

func main() {
	// Deklarasi Variabel
	var g1, g2, g3, g4 string
	var berhasil bool
	
	berhasil = true
	// Perulangan untuk meminta input warna dari pengguna dalam 5 percobaan dan memeriksa apakah input tersebut sesuai dengan warna yang ditentukan 
	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan", i, ": ")
		fmt.Scan(&g1, &g2, &g3, &g4)
		if g1 != "merah" && g2 != "kuning" && g3 != "hijau" && g4 != "ungu" {
			berhasil = false
		}
	}
	fmt.Print("Berhasil: ", berhasil)
}
