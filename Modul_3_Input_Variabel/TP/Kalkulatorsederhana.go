package main

import "fmt"

func main () {
	var angka1, angka2 float32
	var operator string

	print("Masukan angka pertama =") //Masukan angka pertama yang akan diinput
	fmt.Scan(&angka1)

	print("Masukan angka kedua =") // Masukan angka kedua yang akan diinput
	fmt.Scan(&angka2)

	println("Masukan Operator yang digunakan = +, -, /, *, % ")
	fmt.Scan(&operator)

	switch operator {
	case "+" :
		fmt.Print("Hasil dari penjumlahan :",angka1,"+",angka2,"=", angka1 + angka2)
	case "-" :
		fmt.Print("Hasil dari pengurangan :",angka1,"-",angka2,"=", angka1 - angka2)
	case "*" :
		fmt.Print("Hasil dari perkalian :",angka1,"*",angka2,"=", angka1 * angka2)
	case "%" :
	if angka2 !=0 {
		fmt.Print("Hasil dari penjumlahan modulus = " ,  int(angka1) % int(angka2))
	} else {
		fmt.Println("Operasi modulus tidak bisa dioperatorkan karena angka kedua adalah 0")
	}
	case "/" :
	if angka2 !=0 {
		fmt.Print("Hasil dari pembagian :",angka1,"/",angka2,"=", angka1 / angka2)
	} else {
		fmt.Println("Operasi eror karena angka kedua tidak boleh 0")
		}
	}

}
