package main

import "fmt"

func main () {

	var  n, alas, tinggi int

	fmt.Print("Masukan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i ++ {
	fmt.Print("Masukan alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukan tinggi: ")
	fmt.Scan(&tinggi)
	luas := alas * tinggi / 2
	print("luas", luas)
	}
	
}