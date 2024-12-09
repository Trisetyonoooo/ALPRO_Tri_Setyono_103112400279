package main

import "fmt"

func main() {
	var desimal float32

	// Meminta input dari pengguna
	fmt.Print("Masukkan sebuah bilangan desimal: ")
	fmt.Scan(&desimal)

	// Membulatkan bilangan ke atas
	done := int(desimal)
	// Menambahkan 0.1 jika bilangan belum mencapai pembulatan
	if desimal > float32(done) {
		done++
	}
	
	// Menampilkan bilangan hingga mencapai pembulatan ke atas
	for selesai := false; !selesai; {
		// Menampilkan bilangan desimal dengan angka dibelakang koma
		fmt.Printf("%.1f\n", desimal)
		// Menambahkan 0.1
		desimal += 0.1
		// Memeriksa apakah bilangan sudah mencapai pembulatan
		if desimal >= float32(done) {
			// Jika sudah mencapai pembulatan, maka program akan berhenti
			selesai = true
			// Menampilkan bilangan yang sudah mencapai pembulatan
			fmt.Printf("%d\n", done)
		}
	}
}

