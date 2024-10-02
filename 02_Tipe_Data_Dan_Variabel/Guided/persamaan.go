package main

import "fmt"

func main() {
	var x float32
	fmt.Print("Masukkan X: ")
	fmt.Scan(&x)
	var rumus float32 = 2/(x+5) + 5
	fmt.Print("Hasil Persamaan adalah ", rumus)
}