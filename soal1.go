package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga + " " + temp)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga + " " + temp)

	// Program ini meminta 3 inputan dari pengguna, nilai nilai yang diinput akan dimasukan ke dalam variabel satu, dua,tiga
	// Program mencetak output awal dari ketiga variabel tersebut, lalu dilakukan pertukaran/perpindahan nilai sebelumnya
	// Program mencetak output akhir setelah dilakukan proses pertukaran
}