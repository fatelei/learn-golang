package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	top int
	ary [100]string
}

func (s *stack) push(x string) {
	s.ary[s.top] = x
	s.top++
}

func (s *stack) pop() string {
	s.top--
	cur := s.ary[s.top]
	return cur
}

func main() {
	expr := "3+4*5"
	tmp1 := 0

	rst := 0
	s := new(stack)
	symbol := new(stack)

	for _, v := range expr {
		data := string(v)
		switch {
		case "0" <= data && data <= "9":
			s.push(data)
		case "+" == data:
			symbol.push(data)
		case "-" == data:
			symbol.push(data)
		case "/" == data:
			symbol.push(data)
		case "*" == data:
			symbol.push(data)
		}
	}

	for i := symbol.top; i >= 1; i-- {
		tmp := symbol.pop()
		a := s.pop()
		b := s.pop()
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)

		switch tmp {
		case "+":
			tmp1 = x + y
			rst = tmp1 + rst
		case "-":
			tmp1 = y - x
			rst = rst + tmp1
		case "*":
			tmp1 = x * y
			rst = rst + tmp1
		case "/":
			tmp1 = y / x
			rst = rst + tmp1
		}
		s.push(string(tmp1))
	}
	fmt.Println(rst)
}
