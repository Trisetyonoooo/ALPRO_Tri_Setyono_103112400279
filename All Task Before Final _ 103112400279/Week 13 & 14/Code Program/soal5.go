package main

import "fmt"

func main() {

	var x, n, digit int
	var status bool

	fmt.Scan(&x, &n)

	for n > 0 {

		digit = n % 10

		if digit == x {
			status = true
			break
		}

		n /= 10

	}

	fmt.Print(status)

}