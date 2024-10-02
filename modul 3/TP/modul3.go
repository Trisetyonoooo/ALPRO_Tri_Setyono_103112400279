package main

import "fmt"

func main() {
	var (
		sisi int
		hasil int
	)

	fmt.Print("masukan sisi = ")
	fmt.Scan (&sisi)

	hasil = sisi * sisi * sisi

	fmt.Print("Volume kubus = ", hasil)

}