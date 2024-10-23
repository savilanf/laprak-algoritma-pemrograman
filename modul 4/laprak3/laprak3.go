package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64
	fmt.Print("Masukkan Koordinat Titik A (x,y): ")
	fmt.Scan(&ax, &ay)
	fmt.Print("Masukkan Koordinat Titik B (x,y): ")
	fmt.Scan(&bx, &by)
	fmt.Print("Masukkan Koordinat Titik C (x,y): ")
	fmt.Scan(&cx, &cy)
	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ca := math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2))
	fmt.Printf("Panjang sisi terpanjang: %.f\n", math.Max(math.Max(ab, bc), ca))
}