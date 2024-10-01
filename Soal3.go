package main

import "fmt"

func main() {
	var r float64
	var π float64 = 3.14
	var luas float64

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas = π * (r * r) 

	fmt.Print("Luas lingkaran adalah", luas)
}
