package main

import "fmt"

func main() {
    var angka1, angka2, hasil int

    fmt.Print("Masukkan dua buah angka: ")
    fmt.Scan(&angka1, &angka2)

    hasil = 0
    for i := 0; i < angka2; i++ {
        hasil += angka1
    }

    fmt.Printf("Hasil perkalian %d dan %d adalah: %d\n", angka1, angka2, hasil)
}
