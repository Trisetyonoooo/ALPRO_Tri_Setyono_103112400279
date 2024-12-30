package main

import "fmt"

func main () {
	var jumlah, total, suhu, nol int 
	var tertinggi, terendah int 

	for {
		fmt.Scan(&suhu)

		if suhu == 0 {
			nol++
			if nol == 2 {
				break
			}
		} else {
			nol = 0 
			if jumlah == 0 {
				tertinggi = suhu
				terendah = suhu
			} else {
				if suhu > tertinggi {
					tertinggi = suhu
				} else if suhu < terendah {
					terendah = suhu
				}
			}
			total += suhu
			jumlah++
		}
	}

	if jumlah > 0 {
		fmt.Println("Tertinggi:", tertinggi)
		fmt.Println("Terendah:", terendah)
		fmt.Printf("Rata-rata: %.3f\n", float64(total) / float64(jumlah+nol))
	} else {
		fmt.Println("Tidak ada data")
	}
}
