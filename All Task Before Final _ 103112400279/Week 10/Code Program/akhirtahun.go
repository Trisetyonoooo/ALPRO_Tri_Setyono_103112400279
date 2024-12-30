package main 

import "fmt"

func main () {
	var totalbelanja, jumlahdiskon, jumlahcashback, total int 
	var buatkartu, kartu, diskon, cashback bool

	fmt.Scan(&totalbelanja)
	fmt.Scan(&buatkartu)

	kartu = buatkartu

	if totalbelanja >= 100000 {
		diskon = true
		jumlahdiskon = totalbelanja * 10 / 100
	}
	
	if totalbelanja >= 200000 && buatkartu {
		cashback = true
		jumlahcashback = 75000
	}

	total = totalbelanja - jumlahdiskon - jumlahcashback

	fmt.Println("kartu:", kartu)
	fmt.Println("diskon:", diskon)
	fmt.Println("cashback:", cashback)
	fmt.Println("Rp:", total)
}
