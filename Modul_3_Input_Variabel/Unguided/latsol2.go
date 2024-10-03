package main

import "fmt"

func main() {

	const phi = 3.1415926535 // mendefiniskan phi

	var r float64

	fmt.Print("Masukan jari jari =") // input jari jari
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * phi * r * r * r // Menghitung volume bola
	luas := 4 * phi * r * r // Menghitung Luas Kulit bola
	
	//Menampilkan hasil dari operasi volume dan luas
	fmt.Printf("Bola dengan jari jari %.0f, memiliki volume %.4f dan luas kulit bola %.4f\n",r, volume, luas)

	
}