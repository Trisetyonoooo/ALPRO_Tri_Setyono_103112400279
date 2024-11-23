package main

import "fmt"

func main() {
    // Variabel
    var bilangan int
	var hasil int
    // Meminta user memasukan sebuah bilangan
    fmt.Print("Masukan sebuah bilangan: ")
    fmt.Scan(&bilangan)

    // Switch case untuk menentukan dan melakukan operasi bilangan berdasarkan inputnya
    switch {
    case bilangan%2 != 0 && bilangan%5 != 0 && bilangan%10 != 0:
        // Bilangan ganjil (termasuk pengecualian bilangan kelipatan 5 dan 10)
        hasil = bilangan + (bilangan + 1)
        fmt.Printf("Kategori: Bilangan Ganjil\n")
        fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
    case bilangan%2 == 0 && bilangan%5 != 0 && bilangan%10 != 0:
        // Bilangan Genap (termasuk pengecualian bilangan kelipatan 5 dan 10)
        hasil = bilangan * (bilangan + 1)
        fmt.Printf("Kategori: Bilangan Genap\n")
        fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)
    case bilangan%5 == 0 && bilangan%10 != 0:
        // Bilangan Kelipatan 5 (bukan kelipatan 10)
        hasil = bilangan * bilangan
        fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
        fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", bilangan, hasil)
    case bilangan%10 == 0:
        // Bilangan Kelipatan 10
        hasil = bilangan / 10
        fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
        fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)
    default:
        fmt.Println("Tidak ada kategori yang sesuai")
    }
}
