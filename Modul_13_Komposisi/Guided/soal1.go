package main

import "fmt"

func main () {
	var n, i int

	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		if i % 2 != 0 {
				fmt.Print(i, " ")
			}
		}
	}
