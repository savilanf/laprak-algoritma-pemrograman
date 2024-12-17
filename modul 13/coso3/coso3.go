package main

import "fmt"

func main() {
	var x, y int
	var done bool
	fmt.Scan(&x, &y)

	for done = false; !done; {
		x = x - y
		fmt.Println(x)
		done = x <= 0
	}
	fmt.Println(x == 0)
}