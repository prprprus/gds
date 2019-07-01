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

	// case1: get value by index
	v, err := list.Get(3)
	if err != nil && v.(int) != 4 {
		t.Error("case2 error: get value by index")
	}
}

func TestRemove(t *testing.T) {
	list := New(1, 3, 5)

	// case1: remove value by index
	err := list.Remove(1)
	v1, _ := list.Get(0)
	v2, _ := list.Get(1)
	if err != nil && v1.(int) != 1 && v2.(int) != 5 {
		t.Error("case2 error: remove value by index")
	}
	if list.size != 2 {
		t.Error("case2 error: remove value by index, size error")
	}
}

func TestSwap(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: swap value by index
	list.Swap(1, 4)
	v1, _ := list.Get(0)
	v2, _ := list.Get(4)
	if v1.(int) != 4 && v2.(int) != 2 {
		t.Error("case1 error: swap value by index")
	}
}

func TestInsert(t *testing.T) {
	list := New(1, 2, 3)

	// case1: insert a value to mediumn
	// after insert: [1, 2, 88, 3]
	list.Insert(1, 88)
	v1, _ := list.Get(2)
	if v1.(int) != 88 {
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
	// after insert: [1, 2, 88, 3, 11, 12, 56, 66, 83, 81, 80]
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

func TestSet(t *testing.T) {
	list := New(1, 2, 3)

	// case1: set value by index
	list.Set(0, 11)
	list.Set(1, 22)
	list.Set(2, 33)
	v1, _ := list.Get(0)
	v2, _ := list.Get(0)
	v3, _ := list.Get(0)
	if v1.(int) != 11 && v2.(int) != 22 && v3.(int) != 33 {
		t.Error("case1 error: set value by index")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	list := New()

	// case1: list is empty
	if !list.Empty() {
		t.Error("case1 error: list is empty")
	}

	// case2: list is not empty
	list.Append(1, 2, 3)
	if list.Empty() {
		t.Error("case2 error: list is not empty")
	}
}

func TestSize(t *testing.T) {
	list := New(1, 2, 3)

	// case1: size of list
	if list.Size() != 3 {
		t.Error("case1 error: size of list")
	}
}
func TestClear(t *testing.T) {
	list := New(1, 2, 3, 4)

	// case1: clear list
	list.Clear()
	if list.size != 0 && list.frist != nil && list.last != nil {
		t.Error("case1 error: clear list")
	}
}

func TestValues(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: get values
	result := []interface{}{1, 2, 3, 4, 5}
	values := list.Values()
	for i := range result {
		if result[i] != values[i] {
			t.Error("case1 error: get values")
		}
	}
}

// Serialization Interface

func TestToJSON(t *testing.T) {
	list := New(1, 2, 3, 4)

	// case1: list to json
	_, err := list.ToJSON()
	if err != nil {
		t.Error("case1 error: list to json")
	}
}

func TestFromJSON(t *testing.T) {
	list := New(1, 2, 3, 4)

	// case1: list from json
	json, _ := list.ToJSON()
	err := list.FromJSON(json)
	if err != nil {
		t.Error("case1 error: list from json")
	}
}

// Iterator Interface

func TestIterator(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)

	// case1: new iterator
	iterator := list.Iterator()
	if iterator.index != -1 && iterator.element != nil {
		t.Error("case1 error: new iterator")
	}
}

func TestNext(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)
	iterator := list.Iterator()

	// case1: traverse iterator
	for iterator.Next() {
	}
	if iterator.index != 5 && iterator.element != nil {
		t.Error("case1 error: traverse iterator")
	}
}

func TestValue(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)
	iterator := list.Iterator()

	// case1: get value from iterator
	v := 1
	for iterator.Next() {
		if iterator.Value().(int) != v {
			t.Error("case1 error: get value from iterator")
		}
		v++
	}
}

func TestIndex(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)
	iterator := list.Iterator()

	// case1: get index from iterator
	i := 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case1 error: get value from iterator")
		}
		i++
	}
}

func TestBegin(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)
	iterator := list.Iterator()

	// case1: reset iterator
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 && iterator.element != nil {
		t.Error("case1 error: reset iterator")
	}
}
