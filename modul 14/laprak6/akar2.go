package main

import "fmt"

func main() {
	var n int
	var akar float64 = 1.0
	fmt.Scan(&n)

	for k := 0; k <= n; k++ {
		pembilang := float64((4*k + 2) * (4*k + 2))
		penyebut := float64((4*k + 1) * (4*k + 3))
		akar *= pembilang / penyebut
	}
	fmt.Printf("Nilai akar 2 = %.10f \n", akar)
}