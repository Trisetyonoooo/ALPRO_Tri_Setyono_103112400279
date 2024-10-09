package main

import "fmt"

func main() {
	var inputan, digit1, digit2, digit3  int

	fmt.Print("Masukan bilangan 3 digit :")
	fmt.Scan(&inputan)
	
	digit1 = inputan / 100				// Mengambil digit  pertama
	digit2 = (inputan % 100) / 10		// Mengambil digit kedua
	digit3 = inputan % 10				// Mengambil digit ketiga

	fmt.Println(digit1 <= digit2 && digit2 <= digit3)
}