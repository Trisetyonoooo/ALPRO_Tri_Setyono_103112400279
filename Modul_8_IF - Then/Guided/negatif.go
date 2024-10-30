package main

import "fmt"

func main() {
    var bilangan int
    var teks bool

    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)

    if bilangan < 0 && bilangan%2 == 0 {
        teks = true
    } else {
        teks = false
    }

    fmt.Println(teks)
}
