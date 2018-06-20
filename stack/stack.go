package stack

type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack {
		arr: make([]int, 0, 4),
	}
}

func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() (int, bool) {
	if s.Empty() {
		return 0, false
	}

	l := s.Len()
	val := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return val, true
}

func (s *Stack) Top() (int, bool) {
	if s.Empty() {
		return 0, false
	}

	return s.arr[s.Len() - 1], true
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Len() int {
	return len(s.arr)
}
