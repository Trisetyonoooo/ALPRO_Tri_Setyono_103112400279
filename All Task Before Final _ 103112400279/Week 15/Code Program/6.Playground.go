package main 

import "fmt"

func main() {
	var tarif, potongan, diskon, tarif_awal float64
	var durasi, kelebihan int
	var member string 

	fmt.Scan(&member, &durasi)

	if member == "Gold" {
		diskon = 0.5
	} else if member == "Silver" {
		diskon = 0.25
	} else {
		diskon = 0
	}
	if durasi <= 2 {
		tarif_awal =  float64(durasi) * 65000
	} else {
		kelebihan = durasi - 2
		tarif_awal = 2 * 65000 + float64(kelebihan)*20000
	}
	potongan = diskon * tarif_awal
	tarif = tarif_awal - potongan
	fmt.Println("IDR", tarif)
}