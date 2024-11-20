package main

import "fmt"

func main () {
	var jeniskendaraan string
	var durasiparkir, tarif int

	fmt.Print("Masukan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jeniskendaraan)
	fmt.Print("Masukan durasi parkir (dalam jam): ")
	fmt.Scan(&durasiparkir)

	switch {
	case jeniskendaraan == "motor" && durasiparkir >= 1 && durasiparkir <= 2 :
		tarif = 7000
	case jeniskendaraan == "motor" && durasiparkir > 2 :
		tarif = 9000
	case jeniskendaraan == "mobil" && durasiparkir >= 1 && durasiparkir <= 2 :
		tarif = 15000
	case jeniskendaraan == "mobil" && durasiparkir > 2 :
		tarif = 20000
	case jeniskendaraan == "truk" && durasiparkir >= 1 && durasiparkir <= 2 :
		tarif = 25000
	case jeniskendaraan == "truk" && durasiparkir > 2 :
		tarif = 35000
	default :
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif parkir : Rp%d\n", tarif)		
}