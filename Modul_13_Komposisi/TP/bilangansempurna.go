package main 

import "fmt"

func main() {
	var bilangan, jumlah int
	// Meminta user memasukan sebuah bilangan
	fmt.Print("Masukan sebuah bilangan: ")
	fmt.Scan(&bilangan)
	// Perulangan untuk menghitung faktor dari bilangan
	for i := 1; i < bilangan; i++ {
		if bilangan % i == 0 {
			jumlah += i
		}
	}
	// Pengecekan bilangan sempurna
	// Jika jumlah faktor dari bilangan sama dengan bilangan, maka bilangan tersebut bilangan sempurna
	if jumlah == bilangan {
		fmt.Println("Ya",bilangan, "adalah bilangan sempurna")
	} else {
		fmt.Println("Tidak", bilangan, "adalah bukan bilangan sempurna")
	}
}
