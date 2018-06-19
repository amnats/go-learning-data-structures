package list

import (
	"errors"
	"fmt"
)

// double linked list
type List struct {
	head     *node
	tail     *node
	len      int
	iterator *node
}

type node struct {
	next *node
	prev *node
	val  int
}

func New() List {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return List{head: head, tail: tail, iterator: head}
}

func (list *List) Insert(pos int, val int) error {
	if pos < 0 || pos > list.len {
		return errors.New(fmt.Sprintf("insert: cannot insert into list of %v elements to position [%v]", list.len, pos))
	}

	node := &node{val: val}

	if list.len == 0 {
		node.prev = list.head
		node.next = list.tail
		list.head.next = node
		list.tail.prev = node
		list.len++
		return nil
	}

	posNode, err := list.getNode(pos)
	if err != nil {
		return err
	}

	leftNode := posNode.prev
	node.prev = leftNode
	node.next = posNode
	posNode.prev = node
	leftNode.next = node
	list.len++

	return nil
}

func (list *List) PushFront(arr ...int) error {
	for _, val := range arr {
		err := list.Insert(0, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func (list *List) Append(arr ...int) error {
	for _, val := range arr {
		err := list.Insert(list.len, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func (list *List) Erase(pos int) (int, error) {
	if pos < 0 || pos >= list.len {
		return 0, errors.New(fmt.Sprintf("cannot remove element in the list of %v elements to position [%v]", list.len, pos))
	}

	posNode, err := list.getNode(pos)
	if err != nil {
		return 0, err
	}

	oldVal := posNode.val
	leftNode := posNode.prev
	rightNode := posNode.next
	leftNode.next = rightNode
	rightNode.prev = leftNode
	posNode.next = nil
	posNode.prev = nil
	list.len--

	return oldVal, nil
}

func (list *List) PopFront() (int, error) {
	return list.Erase(0)
}

func (list *List) PopBack() (int, error) {
	return list.Erase(list.len - 1)
}

func (list *List) Find(goal int) (*node, bool) {
	node := list.head.next
	for i := 0; i < list.len; i++ {
		if node.val == goal {
			return node, true
		}
		node = node.next
	}

	return nil, false
}

func (list *List) Len() int {
	return list.len
}

func (list *List) Empty() bool {
	return list.len == 0
}

func (list *List) getNode(pos int) (*node, error) {
	if pos > list.len || pos < 0 {
		return nil, errors.New(fmt.Sprintf("cannot get element in the list of %v elements from position [%v]", list.len, pos))
	}

	var node *node
	if pos < list.len/2 {
		node = list.head.next

		for i := 0; i < pos; i++ {
			node = node.next
		}
	} else {
		node = list.tail.prev

		for i := list.len - 1; i > pos; i-- {
			node = node.prev
		}
	}

	return node, nil
}

func (list *List) Get(pos int) (int, error) {
	node, err := list.getNode(pos)
	if err != nil {
		return 0, err
	}

	return node.val, nil
}
