package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const ukuranTetesan = 0.0001

	var banyakTetesan, daerahA, daerahB, daerahC, daerahD int
	var curahA, curahB, curahC, curahD float64

	fmt.Scan(&banyakTetesan)

	for i := 0; i < banyakTetesan; i++ {
		x := rand.Float64()
		y := rand.Float64()
		
		if x < 0.5 && y < 0.5 {
			daerahA++
		} else if x >= 0.5 && y < 0.5 {
			daerahB++
		} else if x < 0.5 && y >= 0.5 {
			daerahC++
		} else {
			daerahD++
		}
	}

	curahA = float64(daerahA) * ukuranTetesan
	curahB = float64(daerahB) * ukuranTetesan
	curahC = float64(daerahC) * ukuranTetesan
	curahD = float64(daerahD) * ukuranTetesan

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahD)
}