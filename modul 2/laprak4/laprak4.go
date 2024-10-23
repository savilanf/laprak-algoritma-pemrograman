package main

import "fmt"

func main() {
	var C, F float64
	fmt.Println("Masukkan fahrenheit")
	fmt.Scan(&F)
	C = (F - 32) * 5 / 9
	fmt.Println("hasilnya adalah", C)
}