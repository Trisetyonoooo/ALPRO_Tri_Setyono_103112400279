package main

import "fmt"

func main() {
	var jam, menit, detik , input int

	fmt.Print("Masukan dalam satuan detik :")
	fmt.Scan(&input)

	jam = input / 3600
	menit = (input % 3600) / 60
	detik = input % 60

	fmt.Println( jam,"jam", menit, "menit", detik, "detik")

}