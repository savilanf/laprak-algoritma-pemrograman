package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luaslingkaran float64
	const phi = 3.14
	fmt.Println("Masukkan jari-jari")
	fmt.Scan(&r)
	luaslingkaran = phi * math.Pow(r, 2)

	fmt.Println("hasilnya adalah", luaslingkaran)
}