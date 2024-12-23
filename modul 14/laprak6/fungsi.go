package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64
	fmt.Scan(&k)

	fk := math.Pow((4 * k + 2), 2) / ((4 * k + 1) * (4 * k + 3))
	fmt.Printf("Nilai f(K) = %.10f\n",  fk)
}