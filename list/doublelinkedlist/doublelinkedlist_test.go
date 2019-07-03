package doublelinkedlist

import "testing"

func TestNew(t *testing.T) {
	// case1: create a new list with no element
	list := New()
	if list.first != nil || list.last != nil || list.size != 0 {
		t.Error("case1 error: create a new list with no element")
	}

	// case2: create a new list with one element
	list = New(1)
	if list.first != list.last || list.size != 1 || list.first.value != 1 {
		t.Error("case2 error: create a new list with one element")
	}

	// case3: create a new list with some elements
	list = New(1, 2, 3, 4)
	if list.size != 4 {
		t.Error("case3 error: create a new list with some elements")
	}
	result := []interface{}{1, 2, 3, 4}
	flag := list.first
	for _, v := range result {
		if flag.value != v {
			t.Error("case3 error: create a new list with some elements")
		}
		flag = flag.next
	}
}

func TestAppend(t *testing.T) {
	// case1: the list has no element, append an empty element
	list := New()
	list.Append()
	if list.first != list.last || list.size != 0 {
		t.Error("case1 error: the list has no element, append an empty element")
	}

	// case2: the list has no element, append one element
	list = New()
	list.Append(1)
	if list.first != list.last || list.first.value != 1 || list.size != 1 {
		t.Error("case2 error: the list has no element, append one element")
	}

	// case3: the list has no element, append some elements
	list = New()
	list.Append(1, 2, 3, 4)
	if list.size != 4 {
		t.Error("case3 error: the list has no element, append some elements")
	}
	result := []interface{}{1, 2, 3, 4}
	flag := list.first
	for _, v := range result {
		if flag.value != v {
			t.Error("case3 error: the list has no element, append some elements")
		}
		flag = flag.next
	}
}
