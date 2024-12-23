package main

import "fmt"

func main() {
	var n, total int
	fmt.Scan(&n)

	for i := 1; i <=n; i++ {
		if i % 2 != 0 {
			total++
		}
	}
	fmt.Printf("Terdapat %v bilangan ganjil", total)
}