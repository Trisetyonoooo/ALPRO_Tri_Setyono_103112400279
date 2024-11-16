// 1. Hasilnya akan eror dan tidak sesuai dengan spesifikasi soal karena terdapat penetapan nilai string pada 
//    variabel nam yang sebelumnya bertipe float

// 2. Kesalahan pada Program tersebut adalah :
//      a. Variabel nam yang diubah menjadi string
//      b. Semua kondisi if berdiri sendiri tanpa else if. Akibatnya beberapa kondisi yang benar dapat menimpa hasil sebelumya
//      c. Variabel nmk tidak digunakan. Akibatnya, output program akan selalu kosong
//      d. Program tidak memvalidasi apakah input nam dari 0 - 100.
// Alur program yang seharusnya :
//      1. Input. Terima input nam
//      2. Validasi Input. Apakah nam berada dalam rentang 0 hingga 100. jika tidak, tampilkan kesalahan dan hentikan eksekusi
//      3. Gunakan struktur if-else-if agar tidak terjadi penumpukan hasil yang sebelumnya
//      4. Output. Program akan menampilkan hasil yang sesuai dengan kriteria soal

// 3. Berikut perbaikan program yang benar :

package main

import "fmt"

func main() {
    var nam float64
    var nmk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam < 0 || nam > 100 {
        fmt.Println("Error: Nilai tidak valid. Masukkan nilai antara 0 hingga 100.")
        return
    }

    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }
    fmt.Println("Nilai mata kuliah:", nmk)
}






