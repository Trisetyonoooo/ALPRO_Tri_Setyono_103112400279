package main

import "fmt"

func main() {
	var kapasitastanki, volumeember, totalvolume int

	fmt.Scan(&kapasitastanki)
	totalvolume = 0

	for {
		fmt.Scan(&volumeember)
		totalvolume += volumeember

		if totalvolume >= kapasitastanki {
			fmt.Println("true") 
			break
		} else {
			fmt.Println("false") 
		}
	}
}