package main

import "fmt"

func main() {
	// deklarasikan variabel
	var bmi, tinggibadan, beratbadan float32
	// inputan nilai BMI untuk User
	fmt.Print("Masukan nilai BMI :")			
	fmt.Scan(&bmi)
	// inputan tinggi badan untuk user
	fmt.Print("Masukan tinggi badan :")			
	fmt.Scan(&tinggibadan)
	// rumus menghitung berat badan
	beratbadan = bmi * (tinggibadan * tinggibadan)	
	// menampilkan hasil nilai berat badan
	fmt.Printf("Berat badannya adalah :%.0f", beratbadan)	
}