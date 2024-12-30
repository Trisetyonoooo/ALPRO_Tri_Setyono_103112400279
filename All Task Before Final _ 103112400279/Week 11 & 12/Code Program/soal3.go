package main 

import "fmt"

func main () {
	var x, digit, jumlah int 

	fmt.Scan(&x)

	for x > 0 {
		digit = x % 10
		fmt.Print(digit, " ")
		jumlah += digit
		x = x / 10
	}
	fmt.Println()
	fmt.Println(jumlah)
}