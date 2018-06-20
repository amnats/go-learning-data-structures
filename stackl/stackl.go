package stackl

import "go-learning-data-structures/list"

type Stackl struct {
	l *list.List
}

func NewStackl() *Stackl{
	return &Stackl{
		l: list.New(),
	}
}

func (s *Stackl) Push(val int) {
	err := s.l.Append(val)
	if err != nil {
		panic(err)
	}
}

func (s *Stackl) Top() (int, bool) {
	val, err := s.l.Get(s.Len() - 1)
	if err != nil {
		return 0, false
	}

	return val, true
}

func (s *Stackl) Pop() (int, bool) {
	val, err := s.l.PopBack()
	if err != nil {
		return 0, false
	}

	return val, true
}

func (s *Stackl) Empty() bool {
	return s.l.Empty()
}

func (s *Stackl) Len() int {
	return s.l.Len()
}

