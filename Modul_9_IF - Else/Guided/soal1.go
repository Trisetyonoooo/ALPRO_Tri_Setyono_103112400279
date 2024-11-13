package main 

import "fmt"

func main() {  
	var usia int  
	var kartukeluarga bool 

	fmt.Scan(&usia, &kartukeluarga) 

	if usia >= 17 && kartukeluarga{  
		fmt.Println("bisa membuat KTP")  
	}else{  
		fmt.Println("belum bisa membuat KTP")  
	} 
}  

