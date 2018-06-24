package pqueue

import (
	"testing"
)

func TestPqueue_Insert(t *testing.T) {
	q := New()

	q.Insert(20)
	q.Insert(30)
	q.Insert(10)

	size := Size(q)
	if size != 3 {
		t.Errorf("got: %v\n expected: %v", size, 3)
	}
}

func TestPqueue_PopMax(t *testing.T) {
	q := New()
	q.Insert(20)
	q.Insert(30)
	q.Insert(10)

	max := q.PopMax()

	if max != 30 {
		t.Errorf("got: %v\n expected: %v", max, 30)
	}
	size := Size(q)
	if size != 2 {
		t.Errorf("got: %v\n expected: %v", size, 2)
	}
}
