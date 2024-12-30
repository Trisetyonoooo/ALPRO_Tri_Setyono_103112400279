package main 

import "fmt"

func main () {
	var transaksi, saldo int 
	
	saldo = 0 

	for {
		fmt.Scan(&transaksi)

		if transaksi == 0 {
			break
		} 
		saldo += transaksi
	}

	fmt.Println(saldo)
}