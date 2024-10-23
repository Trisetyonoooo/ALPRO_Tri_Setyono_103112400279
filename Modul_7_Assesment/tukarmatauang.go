package main

import "fmt"

func main () {
	var  dinar, dirham,fals,qirat int
	fmt.Print("Masukan bilangan bulat yang menyatakan dalam qirat: ")
	fmt.Scan(&qirat)

	dinar = qirat / (10*10*6)
	qirat%= 10 * 10 * 6
	dirham = qirat / (10*6)
	qirat%= 10 * 6
	fals = qirat  / 6
	qirat%= 6
	fmt.Println(dinar,dirham,fals,qirat, " ")
}
	

	
