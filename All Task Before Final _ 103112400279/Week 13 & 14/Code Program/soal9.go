package main

import "fmt"

func main() {
    var x, i, j int

    fmt.Scan(&x) 

    for i = 1; i <= x; i++ { 
        for j = 1; j <= 7; j++ { 
            if i == 1 || i == x || j == 1 || j == 7 { 
                fmt.Print(i) 
            } else {
                fmt.Print(" ") 
            }
        }
        fmt.Println() 
    }
}

