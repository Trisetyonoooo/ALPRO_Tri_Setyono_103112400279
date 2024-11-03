package main 

import "fmt"

func main() {
	var huruf string
	// Inputan Untuk Pengguna memasukan sebuah huruf
	fmt.Print("Masukan satu huruf: ")
	fmt.Scan(&huruf)
	// Memeriksa apakah input pengguna adalah huruf vokal, baik dalam huruf besar ("A", "I", "U", "E", "O") 
	// maupun huruf kecil ("a", "i", "u", "e", "o"). Jika iya program akan menampilkan "Huruf Vokal"
	// Jika tidak program akan menampilkan "Huruf Konsonan"
	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" || 
		huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
			fmt.Println("Huruf Vokal")
		} else {
			fmt.Println("Huruf Konsonan")
		}
}

