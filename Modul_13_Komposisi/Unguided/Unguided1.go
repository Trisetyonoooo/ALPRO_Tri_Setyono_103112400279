package main 

import "fmt"

func main() {
	// Deklarasi variabel
	var bilangan, hasil int
	// Meminta user memasukan sebuah bilangan
	fmt.Print("Masukan sebuah bilangan: ")
	fmt.Scan(&bilangan)
	// Perulangan untuk menghitung jumlah bilangan ganjil dari bilangan yang diinput
	for i := 1; i <= bilangan; i++ {
		if i % 2 != 0 {
			hasil++
		}
	}
	fmt.Println("Terdapat", hasil, "bilangan ganjil")
}