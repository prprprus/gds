package singlylinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	// case1
	list := New()
	if list.frist != nil || list.last != nil || list.size != 0 {
		t.Error("empty list initial error")
	}

	// case2
	list = New(1, 2, 3, 4, 5)
	if list.size != 5 {
		t.Error("list initial error")
	}
}

func TestAppend(t *testing.T) {
	// case1
	list := New(1, 2, 3, 4, 5)
	flag := list.frist
	value := 1
	for flag != nil {
		if flag.value != value {
			t.Error("append value error")
		}
		value++
		flag = flag.next
	}
}
