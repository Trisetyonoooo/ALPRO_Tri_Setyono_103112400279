package main 

import "fmt"

func main () {
	
	var i, y, bilangan, total int

	fmt.Scan(&y)
	for i = 0; i < 9; i++ {
		fmt.Scan(&bilangan)
		total += bilangan
	}
	if total >= y * 5{
		fmt.Println("Median bernilai", y)
	} else {
		fmt.Println("Median bernilai 0")
	}
}