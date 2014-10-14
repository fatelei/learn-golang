package main

import "fmt"

// NormalForLoop is normal for loop func
func NormalForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

// GotoLoop is goto func
func GotoLoop() {
	i := 1
Here:
	if i > 10 {
		return
	}
	fmt.Printf("%d\n", i)
	i = i + 1
	goto Here
}

// ArrayLoop is array loop func
func ArrayLoop() {
	ary := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range ary {
		fmt.Printf("%d\n", v)
	}
}

func main() {
	NormalForLoop()
	GotoLoop()
	ArrayLoop()
}
