package main 

import "fmt"

func main() {
	var n int

	for selesai := false; !selesai; {
		fmt.Print("Masukan angka: ")
		fmt.Scan(&n)
		selesai = (n > 0)
	}
	fmt.Println(n, "adalah bilangan bulat positif")
}
