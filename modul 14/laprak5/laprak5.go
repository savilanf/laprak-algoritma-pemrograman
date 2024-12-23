package main

import "fmt"

func main() {
	var kantong1, kantong2, totalBerat, selisih float64
	var oleng bool

	for totalBerat <= 150 && (kantong1 >= 0 && kantong2 >= 0) {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)
		if kantong2 >= kantong1 {
			selisih = kantong2 - kantong1
		} else if kantong1 >= kantong2 {
			selisih = kantong1 - kantong2
		}

		totalBerat = kantong1 + kantong2
		if totalBerat >= 150 {
			break
		}

		oleng = selisih >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}
	fmt.Println("Proses selesai.")
}