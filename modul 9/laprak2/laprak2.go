package main

import "fmt"

func main() {
	var bilangan int
	var hasil string
	fmt.Scan(&bilangan)
	hasil = "bukan"

	if bilangan < 0 && bilangan % 2 == 0 {
		hasil = "genap negatif"
	}
	fmt.Println(hasil)
}