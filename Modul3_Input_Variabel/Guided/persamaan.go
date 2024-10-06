package main

import "fmt"

func main() {
	var fx int 
	var x  int


	fmt.Print("Masukan nilai x = ")
	fmt.Scanln(&x)

	fx = (2/(x + 5) + 5)


	fmt.Print(fx)
}