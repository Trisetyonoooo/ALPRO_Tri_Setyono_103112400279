package main

import "fmt"

func main () {
	var a,b int

	fmt.Print("Masukan dua bilangan bulat: ")
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		fmt.Print(i)
	}
}