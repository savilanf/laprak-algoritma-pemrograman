package main

import "fmt"

func main() {
	var tipeKendaraan string
	var durasi, tarif int
	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&tipeKendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch {
	case tipeKendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case tipeKendaraan == "Motor" && durasi > 2:
		tarif = 9000
	case tipeKendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case tipeKendaraan == "Mobil" && durasi > 2:
		tarif = 20000
	case tipeKendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case tipeKendaraan == "Truk " && durasi > 2:
		tarif = 35000
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif Parkir : Rp %d", tarif)
}