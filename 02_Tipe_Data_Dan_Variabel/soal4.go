package main

import "fmt"

func main() {
	var ( celcius,fahrenheit int
	)

	fmt.Println("Masukan suhu dalam Fahrenheit:")
	fmt.Scanln(&fahrenheit)

	celcius = (fahrenheit - 32) * 5/9

	fmt.Println("Hasil dari konversi suhu Fahrenheit ke Celcius adalah", celcius)


}