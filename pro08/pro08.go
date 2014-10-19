package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	top int
	ary [10]int
}

func (s *stack) push(v int) {
	if s.top+1 > 9 {
		return
	}
	s.ary[s.top] = v
	s.top++
}

func (s *stack) pop() (v int) {
	v = s.ary[s.top]
	s.top--
	return
}

func (s stack) String() string {
	var str string

	for i := 0; i < s.top; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.ary[i]) + "]"
	}
	return str
}

func main() {
	s := new(stack)
	s.push(10)
	s.push(12)

	cur := s.pop()
	fmt.Println(cur)
	fmt.Printf("%v\n", s)
}
