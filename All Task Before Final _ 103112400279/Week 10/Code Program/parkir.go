package main 

import "fmt"

func main () {
	var h1, h2, m1, m2, total int 
	var jam, menit, waktuawal, waktuakhir int

	fmt.Scan(&h1, &m1, &h2, &m2)

	waktuawal = h1 * 60 + m1
	waktuakhir = h2 * 60 + m2

	if waktuakhir < waktuawal {
		waktuakhir += 12 * 60
	}
	total = waktuakhir - waktuawal

	jam = total / 60
	menit = total % 60

	fmt.Println(jam, "jam", menit, "menit")
}