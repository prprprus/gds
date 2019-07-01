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

// List Interface

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

	// case3: append values to the front of the not empty list
	list = New(6, 7, 8, 9)
	list.PreAppend(1, 2, 3, 4, 5)
	flag = list.frist
	value = 1
	for flag != nil {
		if flag.value != value {
			t.Error("case3 error: append values to the front of the not empty list")
		}
		value++
		flag = flag.next
	}

	// case4: append a value to the front of the not empty list
	list = New(4, 5, 6)
	list.PreAppend(3)
	flag = list.frist
	value = 3
	for flag != nil {
		if flag.value != value {
			t.Error("case4 error: append a value to the front of the not empty list")
		}
		value++
		flag = flag.next
	}

}

func TestIndexInRange(t *testing.T) {
	list := New(1, 2, 3)

	// case1: index in the range
	if !list.indexInRange(0) {
		t.Log("case1 error: index in the range")
	}

	// case2: index out of the range
	if list.indexInRange(5) {
		t.Error("case2 error: index out of range")
	}
}

func TestGet(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: get value
	v, err := list.Get(3)
	if err != nil && v.(int) != 4 {
		t.Error("case2 error: get value")
	}
}

func TestRemove(t *testing.T) {
	list := New(1, 3, 5)

	// case1: remove value
	err := list.Remove(1)
	value1, _ := list.Get(0)
	value2, _ := list.Get(1)
	if err != nil && value1.(int) != 1 && value2.(int) != 5 {
		t.Error("case2 error: remove value")
	}
	if list.size != 2 {
		t.Error("case2 error: remove value, size error")
	}
}

func TestSwap(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: swap value by index
	list.Swap(1, 4)
	value1, _ := list.Get(0)
	value2, _ := list.Get(4)
	if value1.(int) != 4 && value2.(int) != 2 {
		t.Error("case1 error: swap value by index")
	}
}

func TestInsert(t *testing.T) {
	list := New(1, 2, 3)

	// case1: insert a value to mediumn
	// after insert: [1, 2, 88, 3]
	list.Insert(1, 88)
	value1, _ := list.Get(2)
	if value1.(int) != 88 {
		t.Error("case1 error: insert a value to mediumn")
	}

	// case2: insert a value to tailer
	// after insert: [1, 2, 88, 3, 66]
	list.Insert(3, 66)
	value3, _ := list.Get(4)
	if value3.(int) != 66 {
		t.Error("case2 error: insert a value to tailer")
	}

	// case3: insert values to mediumn
	// after insert: [1, 2, 88, 3, 11, 12, 56, 66]
	arr := []interface{}{11, 12, 56}
	list.Insert(3, arr...)
	result := []interface{}{1, 2, 88, 3, 11, 12, 56, 66}
	for i, v := range result {
		value, _ := list.Get(i)
		if value != v {
			t.Error("case3 error: insert values to mediumn")
		}
	}

	// case4: insert values to tailer
	arr = []interface{}{83, 81, 80}
	list.Insert(7, arr...)
	result = []interface{}{1, 2, 88, 3, 11, 12, 56, 66, 83, 81, 80}
	for i, v := range result {
		value, _ := list.Get(i)
		if value != v {
			t.Error("case4 error: insert values to tailer")
		}
	}
}
