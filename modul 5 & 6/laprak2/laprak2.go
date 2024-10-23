package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, tinggi float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Masukkan jejari dan tinggi: ")
		fmt.Scan(&r, &tinggi)
		volume := (1.0 / 3.0) * math.Pi * math.Pow(r,2) * tinggi
		fmt.Printf("%.14f\n", volume)
	}
}