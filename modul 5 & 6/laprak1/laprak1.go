package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Bilangan Bulat Positif: ")
	fmt.Scan(&n)
	hasil := 0
	for i := 1; i <= n; i++ {
		hasil += i
	}
	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah: %d\n", n, hasil)
}