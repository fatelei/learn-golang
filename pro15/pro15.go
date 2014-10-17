package pro15

type Stack struct {
	top int
	ary [10]int
}

func (s *Stack) push(x int) {
	s.ary[s.top] = x
	s.top++
}

func (s *Stack) pop() int {
	s.top--
	cur := s.ary[s.top]
	return cur
}
