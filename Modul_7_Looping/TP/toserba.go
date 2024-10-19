package main

import "fmt"

func main () {
	var i, n, totalpoin int

	fmt.Print("Masukan jumlah barang yang dibeli: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i ++ {
		if i <= 5 {
			// setiap barang sampai barang kelima akan mendapatkan 10 poin
			totalpoin += 10
		} else {
			// Ketika setelah barang kelima maka barang selanjutnya akan mendapat tambahan 5 poin
			totalpoin += 15
		}
	}
	fmt.Println("Total poin yang didapat:", totalpoin, "poin")
}