package main 

import "fmt"

func main () {
	var x rune
	var huruf, vokal bool

	fmt.Scanf("%c", &x)

	huruf = (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')

	vokal = x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o' ||
		x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O'

	if huruf && vokal{
		fmt.Println("Bukan konsonan")
	} else {
		fmt.Println("Konsonan")
	}
}