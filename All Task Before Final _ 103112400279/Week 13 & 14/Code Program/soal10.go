package main

import "fmt"

func main() {
	var x, i, j int
	fmt.Scan(&x)
	for i = 0; i < x; i++ { 
		for j = 0; j < x; j++ { 
			if i == j || i+j == x-1 {
				fmt.Print(j + 1) 
			} else {
				fmt.Print(" ") 
			}
		}
		fmt.Println()
	}
}
