package main

import "fmt"

func main() {
    // Deklarasi variabel
    var n, m, x, y int

    fmt.Scan(&n, &m, &x, &y)
    
    cangkir := 0
    
    // Perulangan untuk menghitung banyak cangkir yang bisa dibuat
    for n >= x && m >= y {
        n -= x  // Kurangi gula untuk satu cangkir
        m -= y  // Kurangi kopi untuk satu cangkir
        cangkir++  // Tambah jumlah cangkir yang bisa dibuat
    }
    
    // Output jumlah cangkir kopi yang bisa dibuat
    fmt.Println(cangkir)
}
