package main

import "fmt"

func main() {
	var orang, motor int
	fmt.Scan(&orang)
	motor = orang / 2

	if orang % 2 != 0 {
		motor = (orang + 1) / 2
	}
	fmt.Println(motor)
}