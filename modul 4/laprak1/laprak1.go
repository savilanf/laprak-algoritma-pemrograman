package main

import "fmt"

func main() {
	var diskon, totalBelanja, totalAkhir int
	fmt.Print("Masukkan Total Belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Masukkan Diskon (%): ")
	fmt.Scan(&diskon)
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Printf("Total Belanja Akhir setelah diskon: %d\n", totalAkhir)
}