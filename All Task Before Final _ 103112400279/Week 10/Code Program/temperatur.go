package main

import "fmt"

func main () {
	var bil1, bil2, bil3, bil4, bil5 float64
	
	fmt.Scan(&bil1, &bil2, &bil3, &bil4, &bil5)

	if bil1 > bil2 && bil2 > bil3 && bil3 > bil4 && bil4 > bil5 {
		fmt.Println("Stabil naik/ turun")
	} else if bil1 < bil2 && bil2 < bil3 && bil3 < bil4 && bil4 < bil5 {
		fmt.Println("Stabil naik/ turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}