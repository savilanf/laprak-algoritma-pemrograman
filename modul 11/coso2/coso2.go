package main

import "fmt"

func main() {
	var namaTanaman string
	fmt.Scan(&namaTanaman)

	switch namaTanaman {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tidak Asli Indonesia")
	default:
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	}
}