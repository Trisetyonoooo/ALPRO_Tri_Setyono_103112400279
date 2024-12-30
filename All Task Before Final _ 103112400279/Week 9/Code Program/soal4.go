package main

import "fmt"

func main () {
	var bilangan int

	fmt.Scan(&bilangan)

	if bilangan %3 == 0 && bilangan %5 == 0 {
		fmt.Println("Kelipatan 3 dan Kelipatan 5")
	} else if bilangan %3 == 0 {
		fmt.Println("Kelipatan 3")
	} else if bilangan %5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}