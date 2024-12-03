package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasi, tarif  int
	fmt.Scan(&jenisKendaraan, &durasi)
	
	switch {
	case jenisKendaraan == "motor" && durasi <= 1:
		tarif = 2000
	case jenisKendaraan == "motor" && durasi > 1:
		tarif = durasi * 2000
	case jenisKendaraan == "mobil" && durasi <= 1:
		tarif = 5000
	case jenisKendaraan == "mobil" && durasi > 1:
		tarif = durasi * 5000
	case jenisKendaraan == "truk" && durasi <= 1:
		tarif = 8000
	case jenisKendaraan == "truk" && durasi > 1:
		tarif = durasi * 8000
	}
	fmt.Printf("Rp %v", tarif)
}