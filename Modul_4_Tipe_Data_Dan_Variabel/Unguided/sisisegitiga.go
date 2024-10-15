package main

import (
	"fmt"
	"math"
)

func main() {
	var(
		Ax, Bx, Cx float64
		Ay, By, Cy float64
		sisiterpanjang float64
		AB, BC, CA float64
	)
	// Inputan untuk memasukan koordinat titik x dan y
	fmt.Print("Masukan titik Ax dan Ay: ")
	fmt.Scan(&Ax, &Ay)
	fmt.Print("Masukan titik Bx dan By: ")
	fmt.Scan(&Bx, &By)
	fmt.Print("Masukan titik Cx dan Cy: ")
	fmt.Scan(&Cx, &Cy)
	// MenGhitung panjang sisi segitiga
	AB = math.Sqrt(math.Pow((Bx-Ax), 2) + math.Pow((By-Ay), 2))
	BC = math.Sqrt(math.Pow((Cx-Bx), 2) + math.Pow((Cy-By), 2))
	CA = math.Sqrt(math.Pow((Ax-Cx), 2) + math.Pow((Ay-Cy), 2))
	// Menentukan sisi terpanjang
	if AB > sisiterpanjang {
		sisiterpanjang = AB
	}
	if BC > sisiterpanjang {
		sisiterpanjang = BC
	}
	if CA > sisiterpanjang {
		sisiterpanjang = CA
	}
	//Menampilkan hasil sisi terpanjang dengan 2 angka dibelakang koma
	fmt.Printf("%.2f\n", sisiterpanjang)	
}