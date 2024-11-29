package main	

import "fmt"	

func main() {
	// Variabel
	var n, digit int
	// Input untuk memasukan bilangan bulat
	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&n)
	// Perulangan untuk memecah bilangan bulat menjadi digit 
	for n > 0 {
		// mengambil digit terakhir menggunakan modulus 
		digit = n % 10
		fmt.Println(digit)
		// menghilangkan digit terakhir dengan pembagian
		n = n / 10
	}
}
