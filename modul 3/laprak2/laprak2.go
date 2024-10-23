package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luasbola, volumebola float64
	fmt.Print("Jejari: ")
	fmt.Scan(&r)
	luasbola = 4 * math.Pi * math.Pow(r, 2)
	volumebola = 4 * math.Pi * math.Pow(r, 3) / 3
	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f\n", r, volumebola, luasbola)
}