package main

import "fmt"

func main () {
	var waktu1, waktu2 int
	var label string

	fmt.Scan(&waktu2)

	switch {
	case waktu2 == 0 :
		waktu1 = 12
		label = "AM"
	case waktu2 < 12 :
		waktu1 = waktu2
		label = "AM"
	case waktu2 == 12:
		waktu1 = 12
		label = "PM"
	case waktu2 > 12 && waktu2 < 23 :
		waktu1 = waktu2 - 12
		label = "PM"
	default :
		fmt.Print("Tidak termasuk jam")
	}
	fmt.Println(waktu1, label)
}