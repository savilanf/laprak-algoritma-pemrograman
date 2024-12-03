package main

import "fmt"

func main() {
	var bilangan, hasil int
	fmt.Scan(&bilangan)

	switch {
	case bilangan % 10 == 0 && bilangan > 10:
		hasil = bilangan / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %v / 10 = %v", bilangan, hasil)
	case bilangan % 5 == 0 && bilangan > 5:
		hasil = bilangan * bilangan
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %v^2 = %v", bilangan, hasil)
	case bilangan % 2 == 0:
		hasil = bilangan * (bilangan + 1)
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %v * %v = %v", bilangan, bilangan+1, hasil)
	case bilangan % 2 != 0:
		hasil = bilangan + (bilangan + 1)
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %v + %v = %v", bilangan, bilangan+1, hasil)
	}
}