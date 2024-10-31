package main

import "fmt"

func main () {
	var x,y int
	var hasilx, hasily bool
	// Inputan untuk memasukan bilangan pertama dan kedua
	fmt.Print("Masukan bilangan pertama: ")
	fmt.Scan(&x)
	fmt.Print("Masukan bilangan kedua: ")
	fmt.Scan(&y)
	// Memeriksa apakah x adalah faktor dari y, yang artinya jika x habis membagi y tanpa ada sisa
	// Maka hasil akan bernilai true, jika tidak hasil akan bernilai false
	hasilx = y%x == 0
	// Memeriksa apakah y adalah faktor dari x, yang artinya jika y habis membagi x tanpa ada sisa
	// Maka hasil akan bernilai true, jika tidak hasil akan bernilai false
	hasily = x%y == 0
	fmt.Println(hasilx)
	fmt.Println(hasily)


}