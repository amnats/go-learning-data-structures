package stackl

import (
	"testing"
)

func TestStackl_Push(t *testing.T) {
	s := NewStackl()
	s.Push(10)

	l := s.Len()
	if l != 1 {
		t.Errorf("Got: %v\nExpected: %v", l, 1)
	}
}

func TestStackl_Top(t *testing.T) {
	s := NewStackl()
	s.Push(10)

	val, _ := s.Top()
	if val != 10 {
		t.Errorf("Got: %v\nExpected: %v", val, 10)
	}
}

func TestStackl_Pop(t *testing.T) {
	s := NewStackl()
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