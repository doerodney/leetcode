package stack

type Stack struct {
	data []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(s.data) - 1
		element := (s.data)[index]
		s.data = (s.data)[:index]
		return element, true
	}
}
