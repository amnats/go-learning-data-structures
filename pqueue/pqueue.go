package pqueue

import (
	"fmt"
)

type Pqueue struct {
	arr []int
	count int
}

func New() *Pqueue {
	return &Pqueue{arr: make([]int, 1, 1), count: 0}
}

func Size(q *Pqueue) (size int) {
	return len(q.arr) - 1
}

func checkIndex(iarr ...int) {
	for _, i := range iarr {
		if i == 0 {
			panic("index out of range")
		}
	}
}

func left(q *Pqueue, pindex int) (p int, i int) {
	checkIndex(pindex)

	i = pindex*2
	if i > Size(q) {
		return 0, 0
	}
	return q.arr[i], i
}

func right(q *Pqueue, pindex int) (p int, i int) {
	checkIndex(pindex)

	i = pindex*2 + 1
	if i > Size(q) {
		return 0, 0
	}
	return q.arr[i], i
}

func parent(q *Pqueue, chindex int) (p int, i int) {
	checkIndex(chindex)

	i = chindex / 2
	return q.arr[i], i
}

func (q *Pqueue) swap(i int, j int) {
	checkIndex(i, j)

	temp := q.arr[i]
	q.arr[i] = q.arr[j]
	q.arr[j] = temp
}

func (q *Pqueue) siftUp(i int) {
	checkIndex(i)

	for pp, pi := parent(q, i); i > 1 && q.arr[i] > pp; pp, pi = parent(q, i) {
		q.swap(i, pi)
		i = pi
	}
}

func (q *Pqueue) siftDown(i int) {
	checkIndex(i)

	maxIndex := i
	p := q.arr[i]

	if lp, li := left(q, i); li != 0 && lp > p {
		maxIndex = li
	}
	if rp, ri := right(q, i); ri != 0 && rp > p {
		maxIndex = ri
	}
	if i != maxIndex {
		q.swap(i, maxIndex)
		q.siftDown(maxIndex)
	}
}

func (q *Pqueue) Insert(p int) {
	q.arr = append(q.arr, p)
	q.siftUp(Size(q))
}

func (q *Pqueue) PopMax() (p int) {
	p = q.arr[1]
	q.arr[1] = q.arr[Size(q)]
	q.arr = q.arr[:Size(q)]
	q.siftDown(1)
	return p
}

func level(q *Pqueue, i int) (lvl int) {
	checkIndex(i)

	for _, li := left(q, i);  li != 0; _, li = left(q, i) {
		lvl++
		i = li
	}
	return lvl
}

func (q *Pqueue) String() (str string) {
	level := level(q, 1)
	levelWidth := 1
	levelCounter := 1

	for i := 1; i <= Size(q); i++ {
		str += fmt.Sprintf(" %v", q.arr[i])

		if levelCounter == levelWidth {
			str += "\n"
			levelCounter = 0
			levelWidth = levelWidth*2
			level--
		}

		levelCounter++
	}

	return str
}
