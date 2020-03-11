// package problem0020

type Stack struct {
	list []byte
}

func (s *Stack) Push(b byte) {
	s.list = append(s.list, b)
}

func (s *Stack) Pop() byte {
	pop := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return pop
}

func isValid(s string) bool {
	stack := &Stack{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack.Push(')')

		case '{':
			stack.Push('}')

		case '[':
			stack.Push(']')
		default:
			if len(stack.list) == 0 || stack.Pop() != s[i] {
				return false
			}
		}
	}
	return len(stack.list) == 0
}
