package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Temperatur Celcius =") //Inputan pengguna untuk memasukan suhu dalam skala celcius
	fmt.Scan(&celcius)

	Reamur := celcius * 4/5 // Konversi suhu dari skala celcius ke Reamur
	kelvin := celcius + 273 // Konversi suhu dari skala Celcius ke kelvin
	Fahrenheit := (celcius * 9/5) + 32 // Konversi suhu dari skala Celcius ke fahrenheit

	// Menampilkan hasil dari konversi suhu
	fmt.Println("Temperatur Reamur", Reamur)
	fmt.Println("Temperatur Fahrenheit", Fahrenheit)
	fmt.Println("Temperatur Kelvin", kelvin)

}