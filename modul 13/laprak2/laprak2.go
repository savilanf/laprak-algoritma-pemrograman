package main

import "fmt"

func main() {
	var des float64
	var n int
	fmt.Scan(&des)
	n = int(des) + 1

	for done := false; !done; {
		des += 0.1
		fmt.Printf("%.1f\n", des)
		done = des > float64(n) -0.11
	}
	fmt.Print(n)

}