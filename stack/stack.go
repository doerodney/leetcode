package stack

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

