package main

import "fmt" 

func main () {
	var keuntungan1, keuntungan2 float64
	var hasil float64

	fmt.Scan(&keuntungan1,&keuntungan2)

	if keuntungan1 < keuntungan2 {
		hasil = keuntungan2 - keuntungan1
		fmt.Println("Peningkatan sebesar", hasil)
	} else if keuntungan1 > keuntungan2 {
		hasil = keuntungan1 - keuntungan2
		fmt.Println("Penurunan sebesar", hasil)
	} else {
		fmt.Println("Tetap")
	}
}