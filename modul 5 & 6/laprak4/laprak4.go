package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial (n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan Bilangan Bulat non negatif: ")
	fmt.Scan(&n)
	hasil := faktorial(n)
	fmt.Print(hasil)

}