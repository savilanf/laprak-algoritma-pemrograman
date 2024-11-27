package main

import "fmt"

func main() {
	var b, faktor int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	faktor = 0
	fmt.Print("Faktor: ")
	
	for i := 1; i <= b; i++ {
		if b % i == 0 {
			fmt.Printf("%d ", i)
			faktor = faktor + 1
		}
	}
	fmt.Println(" ")

	if faktor == 2 {
		fmt.Print("Prima: true")
	} else {
		fmt.Print("Prima: false")
	}
}