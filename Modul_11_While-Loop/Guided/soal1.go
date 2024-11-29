package main

import "fmt"

func main() {
	var n, j int

	fmt.Print("Masukan nilai n: ")
	fmt.Scan(&n)

	j = n

	for j > 1 {
		fmt.Print(j, " x ")
		j = j - 1 
	}
	fmt.Println(1)
}