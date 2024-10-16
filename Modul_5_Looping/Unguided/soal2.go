package main

import (
	"fmt"
	"math"
)

func main() {

	var i, r, tinggi, n int
	var volume float64
	
	fmt.Print("Masukan n: ")
	fmt.Scan(&n)  

	for i = 1; i <= n; i+= 1{
		fmt.Print("Masukan jari jari alas kerucut dan tinggi kerucut: ")
		fmt.Scan(&r, &tinggi)  
		// Menghitung volume kerucut  V = (1/3) * π * r^2 * tinggi
		volume = (1.0 / 3.0) * math.Pi * float64(r) * float64(r) * float64(tinggi)
		// Menampilkan hasil volume kerucut
		fmt.Println("Volume kerucut: ",volume)  
	}
}
