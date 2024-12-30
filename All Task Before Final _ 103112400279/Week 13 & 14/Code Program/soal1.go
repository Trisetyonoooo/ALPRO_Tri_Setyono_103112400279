package main

import "fmt"

func main() {

	var angka int
	var status bool

	status = true
	fmt.Scan(&angka)
	if angka <= 1 {
		status = false
	} else {
		for i := 2; i < angka; i++ {
			if angka%i == 0 {
				status = false
				break
			}
		}
	}
	fmt.Println(status)
}