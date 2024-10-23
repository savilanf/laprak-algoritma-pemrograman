package main

import "fmt"

func main() {
	var celcius, fahrenheit, reamur, kelvin int
	fmt.Print("Masukkan celcius: ")
	fmt.Scan(&celcius)
	fahrenheit = (celcius * 9 / 5) + 32
	reamur = celcius * 4 / 5
	kelvin = (fahrenheit + 460) * 5 / 9
	fmt.Println("fahrenheit:", fahrenheit)
	fmt.Println("reamur:", reamur) 
	fmt.Println("kelvin:", kelvin)
}