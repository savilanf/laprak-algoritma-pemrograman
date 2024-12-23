package main

import "fmt"

func main() {
	var n, faktor int
	var teks string
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			faktor = faktor + 1
		}
	}

	if faktor == 2 {
		teks = "prima"
	} else {
		teks = "bukan prima"
	}
	fmt.Print(teks)
}