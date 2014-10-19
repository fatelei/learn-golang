package main

import "fmt"

func bubbleSort(ary []int) []int {
	length := len(ary)
	swaped := false

	for {
		swaped = false
		for i := 1; i < length; i++ {
			if ary[i-1] > ary[i] {
				ary[i], ary[i-1] = ary[i-1], ary[i]
				swaped = true
			}
		}
		if !swaped {
			break
		}
	}
	return ary
}

func main() {
	ary := []int{5, 4, 3, 2, 1}
	fmt.Println(ary)
	rst := bubbleSort(ary)
	fmt.Println(rst)
}
