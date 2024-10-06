package main

import "fmt"

func main() {
	var (
		fahrenheit float32
		kelvin float32
	)

	fmt.Print("Masukan suhu fahrenheit =")
	fmt.Scanln(&fahrenheit)

	kelvin = (fahrenheit - 32)*5/9 + 273
	fmt.Println("Suhu Kelvinya adalah =",kelvin,"kelvin")

}