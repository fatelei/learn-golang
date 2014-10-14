package main

import "fmt"

func main() {
	var total float64
	values := []float64{1.1, 2.2}
	total = 0

	for _, v := range values {
		total += v
	}

	fmt.Println(total / float64(len(values)))
}
