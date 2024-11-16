package main

import "fmt"

func main() {
	var berat, digit1, digit2, biayaKg, sisaBiaya, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	// Menghitung total berat dalam kg dan sisa gram
	digit1 = berat / 1000   
	digit2 = berat % 1000   

	// Menghitung biaya per kg
	biayaKg = digit1 * 10000

	// Menentukan biaya untuk sisa berat
	if digit2 >= 500 {
		sisaBiaya = digit2 * 5  // Rp. 5 per gram jika lebih dari atau sama dengan 500 gram
	} else {
		sisaBiaya = digit2 * 15 // Rp. 15 per gram jika kurang dari 500 gram
	}

	// Jika total berat lebih dari 10 kg, sisa berat tidak dikenakan biaya tambahan
	if digit1 > 10 {
		sisaBiaya = 0
	}

	// Menghitung total biaya
	totalBiaya = biayaKg + sisaBiaya

	// Output hasil
	fmt.Printf("Detail berat: %d kg + %d gr\n", digit1, digit2)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, sisaBiaya)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
