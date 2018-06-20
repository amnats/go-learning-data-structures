package stack

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := NewStack()
	s.Push(10)

	l := s.Len()
	if l != 1 {
		t.Errorf("Got: %v\nExpected: %v", l, 1)
	}
}

func TestStack_Top(t *testing.T) {
	s := NewStack()
	s.Push(10)

	val, _ := s.Top()
	if val != 10 {
		t.Errorf("Got: %v\nExpected: %v", val, 10)
	}
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	expected := 30
	for !s.Empty() {
		val, _ := s.Pop()
		if val != expected {
			t.Errorf("Got: %v\nExpected: %v", val, expected)
		}
		expected -= 10
	}
}