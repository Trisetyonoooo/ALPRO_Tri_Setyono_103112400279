package main

import "fmt"

func main() {

	var input string
	var typeA, typeB, typeC int

	for {
		fmt.Scan(&input)

		if input == "A" {
			typeA++
		} else if input == "B" {
			typeB++
		} else if input == "C" {
			typeC++
		} else {
			break
		}
	}

	fmt.Println("Tipe A:", typeA)
	fmt.Println("Tipe B:", typeB)
	fmt.Println("Tipe C:", typeC)
}