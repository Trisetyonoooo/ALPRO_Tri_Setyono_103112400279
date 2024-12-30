package main

import "fmt"

func main() {

	var jDaun, lebar, lTerbesar int

	fmt.Scan(&jDaun)
	lTerbesar = 0
	for i := 1; i <= jDaun; i++ {
		fmt.Scan(&lebar)
		if lebar > lTerbesar {
			lTerbesar = lebar
		}
	}
	fmt.Print(lTerbesar)
}