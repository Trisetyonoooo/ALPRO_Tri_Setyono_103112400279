package main

import "fmt"

func main() {

	var tahun int
	var kabisat int

	fmt.Print("Tahun =")
	fmt.Scanln(&tahun)

	kabisat = tahun % 4 // Mengecek apakah tahun kabisat
	if kabisat == 0 { // Jika tahun habis dibagi 400 atau 4 maka output akan bernilai true
		fmt.Println("Tahun", tahun)
		fmt.Println("Kabisat", true)
	} else { // Ketika tahun tidak habis dibagi 400 atau 4 atau mempunyai nilai sisa maka output bernilai false
		fmt.Println("Tahun :", tahun)
		fmt.Println("Kabisat :", false)
	}


}