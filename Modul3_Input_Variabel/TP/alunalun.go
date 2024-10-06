package main

import (
    "fmt"
)

func main(){
    var sisi, luas, keliling int
  
    sisi = 27 // Panjang sisi alun alun
    keliling = 4 * sisi // rumus menghitung keliling persegi
    luas = sisi * sisi  // rumus menghitung luas persegi

    fmt.Println("Keliling alun alun :",keliling)
    fmt.Println("Luas alun alun :",luas)
} 