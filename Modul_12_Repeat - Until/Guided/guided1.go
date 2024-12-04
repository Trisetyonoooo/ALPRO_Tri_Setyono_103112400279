package main

import "fmt"	

func main() {
	var kata string
	var n int

	fmt.Scan(&kata)
	fmt.Scan(&n)

	count := 0
	for i := false; !i; {
		fmt.Println(kata)
		count++
		i = (count >= n)
	}
}