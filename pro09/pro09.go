package main

import "fmt"

// MultiArgs is func
func MultiArgs(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	MultiArgs(1, 2, 3)
}
