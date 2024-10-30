package main

import "fmt"

func main () {
	var bilangan int

	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bilangan)

	if bilangan > 0 {
		fmt.Println("positif")
	} else {
		fmt.Println("bukan positif")
	}
}