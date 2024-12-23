package main

import "fmt"

func main() {
	var jumlah int
	var bunga, all string

	for {
		fmt.Printf("Bunga %v: ", jumlah+1)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}
		all += bunga + " - "
		jumlah++
	}
	fmt.Println("Pita: ", all)
	fmt.Print("Bunga: ", jumlah)
}