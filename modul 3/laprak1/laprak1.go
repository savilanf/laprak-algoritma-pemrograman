package main

import "fmt"

func main() {
	// fx = 2 / (x + 5) + 5
	var fx, x float64
	fmt.Scan(&fx)
	x = 2 / (fx + 5) + 5
	fmt.Print(x)
}