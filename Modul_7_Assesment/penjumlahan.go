package main

import "fmt"

func main () {
	var i, a,b, sum int

	fmt.Print("Masukan dua bilangan positif: ")
	fmt.Scan(&a,&b)
	sum = 0
	for i = a; i <= b;i ++{
		sum += i
	}
	fmt.Println(sum)
}