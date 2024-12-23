package main

import "fmt"

func main() {
	const warna1 = "merah"
	const warna2 = "kuning"
	const warna3 = "hijau"
	const warna4 = "ungu"

	var total int
	var w1, w2, w3, w4 string
	var hasil bool
 
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %v: ", i)
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 == warna1 && w2 == warna2 && w3 == warna3 && w4 == warna4 {
			total++
		}
	}
	hasil = total == 5
	fmt.Println("BERHASIL:", hasil)
}
