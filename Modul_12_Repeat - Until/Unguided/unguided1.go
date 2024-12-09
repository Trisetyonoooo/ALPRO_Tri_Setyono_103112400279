package main 

import "fmt"
func main() {
	// Variabel
	var bilangan, jumlah int

	// Minta user memasukkan sebuah bilangan 
	fmt.Print("Masukan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	// Perulangan untuk menghitung jumlah digit dari bilangan
	for selesai := false; !selesai; {
		// Mengurangi bilangan dengan membagi 10
		bilangan = bilangan / 10
		// Setiap perulangan akan menghitung 1 digit
		jumlah ++
		// Kondisi berhenti perulangan jika bilangan sudah menjadi 0
		selesai = (bilangan == 0)
	}
	// Menampilkan hasil perhitungan jumlah digit
	fmt.Println(jumlah)
}	

