package main

import "fmt"

func main() {
	// Deklarasi variabel
	var bilangan int
	// Meminta user memasukan sebuah bilangan
	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&bilangan)
	// Jika input bilangan kurang dari sama dengan 1 maka program akan berhenti 
	if bilangan <= 1 {
		fmt.Println("Bukan Prima")
		return
	}
	// Perulangan untuk memeriksa apakah bilangan prima atau bukan
	for i := 2; i < bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println("Bukan Prima")
			return
		}
	}
	fmt.Println("Prima")
}


