package main

import "fmt"

func main() {
	var x rune
	var huruf, vokal bool

	fmt.Scanf("%c", &x)

	// Memeriksa apakah karakter adalah huruf (baik kecil maupun besar)
	huruf = (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')

	// Memeriksa apakah karakter adalah vokal (huruf kecil atau besar)
	vokal = x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o' ||
		x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O'

	if huruf && vokal {
		fmt.Println("vokal")
	} else if huruf {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}
