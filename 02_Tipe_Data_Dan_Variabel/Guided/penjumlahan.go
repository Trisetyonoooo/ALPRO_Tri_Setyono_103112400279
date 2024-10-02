package main

import "fmt"

func main() {
	var (a, 
		b, 
		c, 
		d, 
		e, 
		penjumlahan int

	)
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scan(&a, &b, &c, &d, &e)
	penjumlahan = a + b + c + d + e
	fmt.Print("Hasil Penjumlahan adalah ", a,b,c,d,e, "adalah ", penjumlahan)
}