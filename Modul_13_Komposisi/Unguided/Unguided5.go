package main 

import "fmt"

func main() {
	// Deklarasi variabel
	var kantong1, kantong2 float64
	// Perulangan untuk meminta input berat belanjaan
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)
		// Jika kantong 1 atau kantong 2 yang diinputkan kurang dari 0 maka program akan berhenti
		if kantong1 < 0 ||kantong2 < 0 {
			fmt.Println("Proses Selesai")
			break
		}
		// Jika kantong 1 dan kantong 2 yang diinputkan lebih dari 150 maka program akan berhenti
		if kantong1 + kantong2 > 150 {
			fmt.Println("Program Selesai")
			break
		}
		// Jika selisih dari kedua kantong lebih dari 9 maka program akan menampilkan hasil true
		if kantong1 - kantong2 > 9 || kantong2 - kantong1 > 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		// Jika selisih dari kedua kantong kurang dari 9 maka program akan menampilkan hasil false
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}		
	}

}