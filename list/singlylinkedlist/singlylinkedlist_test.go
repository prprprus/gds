package singlylinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	// case1: new a empty list
	list := New()
	if list.frist != nil || list.last != nil || list.size != 0 {
		t.Error("case1 error: new a empty list")
	}

	// case2: new a not empty list
	list = New(1, 2, 3, 4, 5)
	if list.size != 5 {
		t.Error("case2 error: new a not empty list")
	}
}

func TestAppend(t *testing.T) {
	// case1: new a empty list and append values
	list := New(1, 2, 3, 4, 5)
	flag := list.frist
	value := 1
	for flag != nil {
		if flag.value != value {
			t.Error("case1 error: append to empty list with more than one value")
		}
		value++
		flag = flag.next
	}

	// case2: new a empty list and append a value
	list = New(1)
	if list.frist.value != 1 && list.size == 1 {
		t.Error("case2 error: append to empty list with one value")
	}

	// case3: append values to not empty list
	list = New(1, 2, 3, 4, 5)
	list.Append(6, 7, 8, 9)
	flag = list.frist
	value = 1
	for flag != nil {
		if flag.value != value {
			t.Error("case3 error: append to not empty list with more than on value")
		}
		value++
		flag = flag.next
	}

	// case4: append a value to not empty list
	list = New(1, 2, 3, 4, 5)
	list.Append(9)
	if list.last.value != 9 && list.size == 6 {
		t.Error("case4 error: append to not empty list with one value")
	}
}

func TestPreAppend(t *testing.T) {
	// case1: append values to the front of the empty list
	list := New()
	list.PreAppend(1, 2, 3, 4, 5)
	flag := list.frist
	value := 1
	for flag != nil {
		if flag.value != value {
			t.Error("case1 error: append values to the front of the empty list")
		}
		value++
		flag = flag.next
	}

	// case2: append a value to the front of the empty list
	list = New()
	list.PreAppend(1)
	if list.frist.value != 1 && list.size != 1 {
		t.Error("case2 error: append a value to the front of the empty list")
	}
}
