package main

import "fmt"

// FizzBuzz is a fizz-buzz func
func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		}
	}
}

func main() {
	FizzBuzz()
}
