package main

import "fmt"

func main () {
	// Deklarasi variabel
	var i, n, kuadrat int
	// Meminta input dari pengguna 
	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&n)
	// Melooping dari 1 hinnga n untuk mencetak kuadrat dari setiap bilangan
	for i = 1; i <= n; i++ {
		kuadrat = i * i
		fmt.Print(kuadrat, " ")
	}
}