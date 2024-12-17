package main

import "fmt"

func main() {
	var target, donasi, total int
	fmt.Scan(&target)
	donatur := 0
	
	for done := false; !done; {
		fmt.Scan(&donasi)
		total = total + donasi
		donatur++
		fmt.Printf("Donatur %v: menyumbang %v. Total terkumpul: %v\n", donatur, donasi, total)
		done = total >= target
	}
	fmt.Printf("Target tercapai! Total donasi: %v dari %v donatur.", total, donatur)
}