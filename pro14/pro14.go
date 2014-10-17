package main

import "fmt"

func plusTwo() func(x int) int {
	real := func(x int) int {
		x = x + 2
		return x
	}
	return real
}

func plusX(x int) func(int) int {
	a := func(n int) int {
		n = n + x
		return n
	}
	return a
}

func main() {
	p := plusTwo()
	fmt.Println(p(2))

	px := plusX(5)
	fmt.Println(px(7))
}
