package main

import "fmt"

func main() {
	
	var x, digit1, digitselanjutnya int
	var konsekutif bool
	
	fmt.Scan(&x)

	digit1 = x % 10
	x = x / 10
	konsekutif = true

	for x > 0 {
		digitselanjutnya = x % 10
		if digitselanjutnya != digit1 +1 && digitselanjutnya != digit1-1 {
			konsekutif = false
			break
		}
		digit1 = digitselanjutnya
		x = x / 10
	}

	if konsekutif {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
