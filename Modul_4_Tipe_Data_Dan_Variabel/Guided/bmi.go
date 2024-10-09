package main

import "fmt"

func main() {
	var bb, tb, hasil float32

	fmt.Print("Masukan berat badan :")
	fmt.Scan(&bb)
	fmt.Print("Masukan tinggi badan :")
	fmt.Scan(&tb)

	hasil = bb / (tb * tb)

	fmt.Printf(" Hasil dari nilai BMI adalah : %.2f", hasil)

}
