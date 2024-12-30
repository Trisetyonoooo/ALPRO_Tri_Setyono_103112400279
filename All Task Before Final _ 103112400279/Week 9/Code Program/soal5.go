package main

import "fmt"

func main() {
	var bil1, bil2, bil3, bil4 int
	var terbanyak, tersedikit int

	fmt.Scan(&bil1, &bil2, &bil3, &bil4)

	terbanyak = bil1
	tersedikit = bil1

	if bil2 > terbanyak {
		terbanyak = bil2
	}
	if bil3 > terbanyak {
		terbanyak = bil3
	}
	if bil4 > terbanyak {
		terbanyak = bil4
	}
	if bil2 < tersedikit {
		tersedikit = bil2
	}
	if bil3 < tersedikit {
		tersedikit = bil3
	}
	if bil4 < tersedikit {
		tersedikit = bil4
	}
	fmt.Println("Jumlah gol terbanyak:", terbanyak)
	fmt.Println("Jumlah gol tersedikit:", tersedikit)
}
