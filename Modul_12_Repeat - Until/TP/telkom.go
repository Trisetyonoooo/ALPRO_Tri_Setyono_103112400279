package main		

import "fmt"


func main() {
	// Variabel 
	var kata string
	// Perulangan untuk memasukan kata sampai user memasukan kata telkom
	for {
		// Meminta input dari user
		fmt.Print("Masukan kata: ")
		fmt.Scan(&kata)
		// Perulangan akan berhenti ketika user memasukan kata telkom
		if kata == "telkom" {
			fmt.Println("Program selesai")
			break
		}
		// jika kata yang dimasukan bukan telkom, program akan melanjutkan perulangan dan menampilkan pesan ini
		fmt.Println("Anda mengetik: ", kata)
	}
}