package main

import "fmt"

func main() {
	
	fmt.Println("Berikut adalah bilangan genap dari 1 sampai 50: ")

	for i := 1; i <= 50; i ++ {
		// Mengecek apakah i habis dibagi dua(bilangan genap)
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
