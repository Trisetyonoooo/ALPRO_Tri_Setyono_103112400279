package main

import "fmt"

func main() {
    var a, b, c, max, min int

    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

   if a > b {
	max = a
	min = b
   } else {
	max = b
	min = a
   }
   if max < c {
	max = c
   }
   if min > c {
	min = c
   }
   fmt.Println("Terbesar:", max)
   fmt.Println("Terkecil:", min)
}