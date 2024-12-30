package main

import "fmt"

func main() {
    const correctUsername = "admin"
    const correctPassword = "admin"
    var username, password string
    var gagalLogin int

    for {
        fmt.Scan(&username)
        fmt.Scan(&password)
		if username == correctUsername && password == correctPassword {
            break
        } else {
            gagalLogin++
        }
    }

    fmt.Println(gagalLogin, "Login Berhasil")
}
