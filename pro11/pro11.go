package main

import "fmt"

// Map is func
func Map(f func(int) int, ary ...int) []int {
	rst := make([]int, len(ary))

	for i, v := range ary {
		rst[i] = f(v)
	}
	return rst
}

func main() {
	a := func(x int) int {
		return x * x
	}
	rst := Map(a, 1, 2)
	fmt.Println(rst)
}
