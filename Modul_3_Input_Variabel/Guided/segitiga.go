package main

import "fmt"

func main() {
	var (
		luas, a, t float64
	)

	fmt.Print("Masukan alas segitiga = ")
	fmt.Scan(&a)
	fmt.Print("Masukan alas segitiga = ")
	fmt.Scan(&t)

	luas = 0.5 * a * t

	println("hasil luas segititga = ", luas )

}
