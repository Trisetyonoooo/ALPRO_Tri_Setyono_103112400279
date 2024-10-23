package main

import "fmt"

func main () {
	var i,n int

	fmt.Print("Masukan bilangan bulat positif: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i ++{
		hasil := i * i
		fmt.Println("hasilnya adalah: ", hasil)
	}
}