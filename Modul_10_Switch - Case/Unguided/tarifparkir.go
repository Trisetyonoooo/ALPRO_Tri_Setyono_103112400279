package main

import "fmt"

func main() {
    var jeniskendaraan string
    var durasiparkir, tarif int
	// Inputan untuk pengguna memasukan jenis kendaraan
    fmt.Print("Masukan jenis kendaraan (motor/mobil/truk): ")
    fmt.Scan(&jeniskendaraan)
	// inputan untuk memasukan durasi parkir dalam jam
    fmt.Print("Masukan durasi parkir (dalam jam): ")
    fmt.Scan(&durasiparkir)

    // Jika durasi parkir kurang dari 1 jam, maka durasi tetap dianggap sebagai 1 jam
    switch {
    case durasiparkir < 1:
            durasiparkir = 1
    } 
	// Menggunakan switch case untuk menghitung total biaya parkir sesuai dengan jenis kendaraanya
    switch jeniskendaraan {
    case "motor":
        tarif = 2000 * durasiparkir
    case "mobil":
        tarif = 5000 * durasiparkir
    case "truk":
        tarif = 8000 * durasiparkir
    default:
        fmt.Println("Jenis kendaraan tidak valid")
    }
    // Output total biaya parkir
    fmt.Printf("Tarif parkir: Rp%d\n", tarif)
}
