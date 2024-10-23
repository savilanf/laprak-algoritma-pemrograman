package main

import "fmt"

func main() {
	var b1, b2 int
	fmt.Scan(&b1, &b2)
	hasil := 1
	for i := 0; i < b2; i++ {
		hasil *= b1
	}
	fmt.Print(hasil)
}