package main

import "fmt"

func main() {
	var angka, digit, i int
	fmt.Scan(&angka)
	digit = 1

	for proses := true; proses; {
		digit = digit * 10
		proses = digit < angka
		i++
	}
	fmt.Println(i)
}