package main

import "fmt"

func main() {
	var (
		dolar,
		rupiah,
		konversi int
	)

	dolar = 15000 

	fmt.Scan(&rupiah)

	konversi = rupiah / dolar

	println(konversi)

}