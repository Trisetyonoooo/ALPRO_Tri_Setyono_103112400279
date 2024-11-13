package main

import "fmt"

func main() {
	var bilangan int
	var digit1, digit2, digit3, digit4 int

	fmt.Scan(&bilangan)

	// Memastikan bilangan harus memiliki tepat 4 digit
	if bilangan >= 1000 && bilangan <= 9999 {
		// Memisahkan digit-digit dari bilangan
		digit1 = bilangan / 1000
		digit2 = (bilangan % 1000) / 100
		digit3 = (bilangan % 100) / 10
		digit4 = bilangan % 10

		if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
			fmt.Println("Terurut membesar")
		} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
			fmt.Println("Terurut mengecil")
		} else {
			fmt.Println("Tidak terurut")
		}
	} else {
		// Pesan kesalahan jika bilangan tidak memiliki 4 digit
		fmt.Println("Bilangan harus terdiri dari 4 digit")
	}
}
