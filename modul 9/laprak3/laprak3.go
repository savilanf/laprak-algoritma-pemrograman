package main

import "fmt"

func main() {
	var x, y int
	var hasil bool
	fmt.Scan(&x, &y)
	
	// cek x faktor dari y
	if y % x == 0 {
		hasil = true
	} else {
		hasil = false
	}
	fmt.Println(hasil) 

	// cek y faktor dari x
	if x % y == 0 {
		hasil = true
	} else {
		hasil = false
	}
	fmt.Println(hasil)
}