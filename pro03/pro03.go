package main

import (
	"fmt"
	"unicode/utf8"
)

// PrintTriangle is print triangle func
func PrintTriangle() {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("A")
		}
		fmt.Printf("\n")
	}
}

// BytesCount is func that counts the number of bytes
func BytesCount(s string) {
	count := 0
	blank := " "
	for _, i := range s {
		if string(i) != blank {
			count++
		}
	}
	fmt.Printf("The runes are: %d\n", count)

	fmt.Printf("The bytes count are: %d\n", utf8.RuneCountInString(s))
}

// ReverseWord reverse the word
func ReverseWord() {
	var s []rune
	s = []rune("foobar")

	reverse := make([]rune, len(s))

	for i := 0; i < len(s); i++ {
		reverse[i] = s[len(s)-1-i]
	}

	fmt.Println(string(reverse))
}

func main() {
	PrintTriangle()
	BytesCount("asSASA ddd dsjkdsjs dk")
	BytesCount("asSAabcddd dsjkdsjs dk")
	ReverseWord()
}
