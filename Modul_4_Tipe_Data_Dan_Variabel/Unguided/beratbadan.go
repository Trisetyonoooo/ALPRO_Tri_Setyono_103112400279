package main

import "fmt"

func main() {
	var bmi, tinggibadan, beratbadan float32				// deklarasikan variabel

	fmt.Print("Masukan nilai BMI :")						// inputan nilai BMI untuk User
	fmt.Scan(&bmi)
	fmt.Print("Masukan tinggi badan :")						// inputan tinggi badan untuk user
	fmt.Scan(&tinggibadan)

	beratbadan = bmi * (tinggibadan * tinggibadan)			// rumus menghitung berat badan
	fmt.Printf("Berat badannya adalah :%.0f", beratbadan)	// menampilkan hasil nilai berat badan
}