package main

import "fmt"

func main() {

	var angka int
	fmt.Scan(&angka)
	biner := ""
	for angka > 0 {
		biner = fmt.Sprintf("%d%s", angka%2, biner)
		angka /= 2
	}
	fmt.Println(biner)
}