package main

import "fmt"

func main() {
	var n, jumlah, rerata, i float64

	for {
		fmt.Scan(&n)
		if n == 9999 {
			break
		}
		i = i + 1
		jumlah = jumlah + n
	}
	rerata = jumlah / i
	fmt.Print(rerata)
}