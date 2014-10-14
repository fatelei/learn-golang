package main

import "fmt"

// Average used to calculate the average value
func Average() {
	var total float64
	var length float64
	total = 0.0

	ary := []float64{1.1, 2.2}
	length = float64(len(ary))

	for _, v := range ary {
		total += v
	}

	fmt.Println(total / length)
}

func main() {
	Average()
}
