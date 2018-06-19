package list

import (
	"testing"
)

func TestList_PushFront(t *testing.T) {
	list := New()

	err := list.PushFront(10)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if list.Len() != 1 {
		t.Errorf("len is not %v", list.Len())
		t.FailNow()
	}

	val, err := list.Get(0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if val != 10 {
		t.Errorf("\nPushFront failed\nGot: %v\nExpected: %v", val, 10)
	}
}

func TestList_Append(t *testing.T) {
	list := New()

	err := list.Append(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if list.Len() != 1 {
		t.Errorf("len is not %v", list.Len())
		t.FailNow()
	}

	val, err := list.Get(0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if val != 10 {
		t.Errorf("\nAppend failed\nGot: %v\nExpected: %v", val, 10)
	}
}

func TestList_Insert_ToEmpty(t *testing.T) {
	list := New()

	err := list.Insert(0, 10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	val, err := list.Get(0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if val != 10 {
		t.Errorf("\nGot: %v\nExpected: %v", val, 10)
	}
}

func TestList_Insert_ToMiddle(t *testing.T) {
	list := New()
	list.Append(10)
	list.Append(30)

	err := list.Insert(1, 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	val, err := list.Get(1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if val != 20 {
		t.Errorf("\nGot: %v\nExpected: %v", val, 20)
	}
}

func TestList_Insert_Error(t *testing.T) {
	list := New()

	err := list.Insert(1, 10)
	if err == nil {
		t.Error("There should be an error")
		t.FailNow()
	}
}

func TestList_Find(t *testing.T) {
	list := New()
	err := list.Append(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	node, found := list.Find(10)

	if !found {
		t.Error("not found")
		t.FailNow()
	}

	if node.val != 10 {
		t.Errorf("found: %v/n expected: %v", node.val, 10)
		t.FailNow()
	}
}

func TestList_Erase(t *testing.T) {
	list := New()
	err := list.Append(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestList_PopBack(t *testing.T) {
	list := New()
	err := list.Append(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestList_PopFront(t *testing.T) {
	list := New()
	err := list.Append(10)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
