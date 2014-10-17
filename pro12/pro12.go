package main

import "fmt"

func findMin(ary []int) int {
	min := ary[0]

	for i := 1; i < len(ary); i++ {
		if ary[i] < min {
			min = ary[i]
		}
	}
	return min
}

func findMax(ary []int) int {
	max := ary[0]

	for i := 1; i < len(ary); i++ {
		if ary[i] > max {
			max = ary[i]
		}
	}
	return max
}

func main() {
	y := []int{1, 2, 3, 4, 5}
	min := findMin(y)
	max := findMax(y)
	fmt.Printf("Min: %d\n", min)
	fmt.Printf("Max: %d\n", max)
}
