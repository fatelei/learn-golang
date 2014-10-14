package main

import "fmt"

// OrderAsc used to get the rihjy order
func OrderAsc(x, y int) (a, b int) {
	if x > y {
		a, b = y, x
	} else {
		a, b = x, y
	}
	return
}

func main() {
	a, b := OrderAsc(7, 2)
	fmt.Println(a, b)
}
