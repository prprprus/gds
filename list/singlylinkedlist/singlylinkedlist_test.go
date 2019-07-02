// singlylinkedlist test principle:
// 1. empty list
// 2. only one element
// 3. two element
// 4. handle first element and last element

package singlylinkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New()

	// case1: new a empty list
	if list.first != nil || list.last != nil || list.size != 0 {
		t.Error("case1 error: new a empty list")
	}

	list = New(1, 2, 3, 4, 5)

	// case2: new a not empty list
	if list.size != 5 {
		t.Error("case2 error: new a not empty list")
	}
}

// List Interface

func TestAppend(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: append values to empty list
	flag := list.first
	value := 1
	for flag != nil {
		if flag.value != value {
			t.Error("case1 error: append values to empty list")
		}
		value++
		flag = flag.next
	}

	list = New(1)

	// case2: append one value to empty list
	if list.first.value != 1 && list.size == 1 {
		t.Error("case2 error: append one value to empty list")
	}

	list = New(1, 2, 3, 4, 5)

	// case3: append values to not empty list
	list.Append(6, 7, 8, 9)
	flag = list.first
	value = 1
	for flag != nil {
		if flag.value != value {
			t.Error("case3 error: append values to not empty list")
		}
		value++
		flag = flag.next
	}

	list = New(1, 2, 3, 4, 5)

	// case4: append one value to not empty list
	list.Append(9)
	if list.last.value != 9 && list.size == 6 {
		t.Error("case4 error: append one value to not empty list")
	}
}

func TestPreAppend(t *testing.T) {
	list := New()

	// case1: append values to the front of the empty list
	list.PreAppend(1, 2, 3, 4, 5)
	flag := list.first
	value := 1
	for flag != nil {
		if flag.value != value {
			t.Error("case1 error: append values to the front of the empty list")
		}
		value++
		flag = flag.next
	}

	list = New()

	// case2: append a value to the front of the empty list
	list.PreAppend(1)
	if list.first.value != 1 && list.size != 1 {
		t.Error("case2 error: append a value to the front of the empty list")
	}

	list = New(6, 7, 8, 9)

	// case3: append values to the front of the not empty list
	list.PreAppend(1, 2, 3, 4, 5)
	flag = list.first
	value = 1
	for flag != nil {
		if flag.value != value {
			t.Error("case3 error: append values to the front of the not empty list")
		}
		value++
		flag = flag.next
	}

	list = New(4, 5, 6)

	// case4: append a value to the front of the not empty list
	list.PreAppend(3)
	flag = list.first
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

	// case1: the first index in the range
	if !list.indexInRange(0) {
		t.Log("case1 error: the first index in the range")
	}

	// case2: the medium index in the range
	if !list.indexInRange(1) {
		t.Log("case2 error: the medium index in the range")
	}

	// case3: the last index in the range
	if !list.indexInRange(2) {
		t.Error("case2 error: the last index in the range")
	}

	// case4: the index out of the range
	if list.indexInRange(3) {
		t.Error("case4 error: the index out of the range")
	}
}

func TestGet(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: get medium value by index
	v, err := list.Get(3)
	if err != nil && v.(int) != 4 {
		t.Error("case1 error: get medium value by index")
	}

	// case2: get first value by index
	v, err = list.Get(0)
	if err != nil && v.(int) != 1 {
		t.Error("case2: get first value by index")
	}

	// case3: get last value by index
	v, err = list.Get(4)
	if err != nil && v.(int) != 5 {
		t.Error("case2: get last value by index")
	}
}

func TestRemove(t *testing.T) {
	list := New(1, 3, 4, 5, 6)

	// case1: remove first value by index
	_ = list.Remove(0)
	v, _ := list.Get(0)
	if list.size != 4 && v.(int) != 3 {
		t.Error("case1 error: remove first value by index")
	}

	// case2: remove last value by index
	_ = list.Remove(3)
	v, _ = list.Get(2)
	if list.size != 3 && v.(int) != 5 {
		t.Error("case2 error: remove last value by index")
	}

	// case3: remove medium value by index
	_ = list.Remove(1)
	v, _ = list.Get(1)
	if list.size != 2 && v.(int) != 5 {
		t.Error("case3 error: remove medium value by index")
	}

	list = New()

	// case4: remove value by index from empty list
	err1 := list.Remove(0)
	err2 := list.Remove(3)
	if err1 == nil && err2 == nil {
		t.Error("case4 error: remove value by index from empty list")
	}

	list = New(1)

	// case5: remove value by index from a list of only one element
	_ = list.Remove(0)
	if list.size != 0 {
		t.Error("case5 error: remove value by index from a list of only one element")
	}
}

func TestSwap(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: swap value by index
	// [1, 4, 3, 2, 5]
	_ = list.Swap(1, 3)
	v1, _ := list.Get(1)
	v2, _ := list.Get(3)
	if v1.(int) != 4 || v2.(int) != 2 {
		t.Error("case1 error: swap value by index")
	}

	// case2: swap value by equal index
	_ = list.Swap(2, 2)
	v3, _ := list.Get(2)
	if v3.(int) != 3 {
		t.Error("case2 error: swap value by equal index")
	}

	// case3: swap first value by index
	// [3, 4, 1, 2, 5]
	_ = list.Swap(0, 2)
	v4, _ := list.Get(0)
	v5, _ := list.Get(2)
	if v4.(int) != 3 || v5.(int) != 1 {
		t.Error("case3 error: swap first value by index")
	}

	// case4: swap last value by index
	// [3, 4, 1, 5, 2]
	_ = list.Swap(3, 4)
	v6, _ := list.Get(3)
	v7, _ := list.Get(4)
	if v6.(int) != 5 || v7.(int) != 2 {
		t.Error("case4 error: swap last value by index")
	}
}

func TestInsert(t *testing.T) {
	list := New(1, 2, 3)

	// case1: insert a value after first element
	// [1, 88, 2, 3]
	_ = list.Insert(0, 88)
	v1, _ := list.Get(1)
	if v1.(int) != 88 {
		t.Error("case1 error: insert a value after first element")
	}

	// case2: insert a value after last element
	// [1, 88, 2, 3, 66]
	_ = list.Insert(3, 66)
	v1, _ = list.Get(4)
	if v1.(int) != 66 {
		t.Error("case2 error: insert a value after last element")
	}

	// case3: insert a value after medium element
	// [1, 88, 2, 7, 3, 66]
	_ = list.Insert(2, 7)
	v1, _ = list.Get(3)
	if v1.(int) != 7 {
		t.Error("case3 error: insert a value after medium element")
	}

	// case4: insert values after medium element
	// [1, 88, 2, 7, 11, 12, 56, 3, 66]
	arr := []interface{}{11, 12, 56}
	list.Insert(3, arr...)
	result := []interface{}{1, 88, 2, 7, 11, 12, 56, 3, 66}
	for i, v := range result {
		value, _ := list.Get(i)
		if value != v {
			t.Error("case4 error: insert values after medium element")
		}
	}

	// case5: insert values after first element
	// [1, 5, 32, 70, 88, 2, 7, 11, 12, 56, 3, 66]
	arr = []interface{}{5, 32, 70}
	list.Insert(0, arr...)
	result = []interface{}{1, 5, 32, 70, 88, 2, 7, 11, 12, 56, 3, 66}
	for i, v := range result {
		value, _ := list.Get(i)
		if value != v {
			t.Error("case5 error: insert values after first element")
		}
	}

	// case6: insert values after last element
	// [1, 5, 32, 70, 88, 2, 7, 11, 12, 56, 3, 66, 0, 9, 87, 111]
	arr = []interface{}{0, 9, 87, 111}
	list.Insert(11, arr...)
	result = []interface{}{1, 5, 32, 70, 88, 2, 7, 11, 12, 56, 3, 66, 0, 9, 87, 111}
	for i, v := range result {
		value, _ := list.Get(i)
		if value != v {
			t.Error("case6 error: insert values after last element")
		}
	}

	list = New(1)

	// case7: insert a value to a list of only one element
	_ = list.Insert(0, 1)
	v1, _ = list.Get(0)
	v2, _ := list.Get(1)
	if list.size != 2 || v1.(int) != 1 || v2.(int) != 1 {
		t.Error("case7 error: insert a value to a list of only one element")
	}

	list = New(1)

	// case8: insert values to a list of only one element
	arr = []interface{}{23, 233, 2333}
	_ = list.Insert(0, arr...)
	result = []interface{}{1, 23, 233, 2333}
	for i, v := range result {
		value, _ := list.Get(i)
		if value != v {
			t.Error("case8 error: insert values to a list of only one element")
		}
	}

	list = New()

	// case9: insert a value to a list of empty
	err := list.Insert(0, 1)
	if err == nil {
		t.Error("case9 error: insert a value to a list of empty")
	}

	// case10: insert values to a list of empty
	err = list.Insert(0, 1, 2, 3)
	if err == nil {
		t.Error("case10 error: insert values to a list of empty")
	}
}

func TestSet(t *testing.T) {
	list := New(1, 2, 3)

	// case1: set first value by index
	_ = list.Set(0, 11)
	v, _ := list.Get(0)
	if v.(int) != 11 {
		t.Error("case1 error: set first value by index")
	}

	// case2: set last value by index
	_ = list.Set(2, 33)
	v, _ = list.Get(2)
	if v.(int) != 33 {
		t.Error("case2 error: set last value by index")
	}

	// case3: set medium value by index
	_ = list.Set(1, 22)
	v, _ = list.Get(1)
	if v.(int) != 22 {
		t.Error("case3 error: set last value by index")
	}

	list = New(1)

	// case4: set value by index with a list of only one element
	_ = list.Set(0, -1)
	v, _ = list.Get(0)
	if v.(int) != -1 {
		t.Error("case4 error: set value by index from a list of only one element")
	}

	list = New()

	// case5: set value by index with a list of empty
	err := list.Set(0, 99)
	if err == nil {
		t.Error("case5 error: set value by index with a list of empty")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	list := New()

	// case1: the list is empty
	if !list.Empty() {
		t.Error("case1 error: the list is empty")
	}

	list = New(1)

	// case2: the list has one element
	if list.Empty() {
		t.Error("case2 error: the list has one element")
	}

	list = New(1, 2, 3, 4)

	// case3: the list has some elements
	if list.Empty() {
		t.Error("case3 error: the list has some elements")
	}
}

func TestSize(t *testing.T) {
	list := New(1, 2, 3)

	// case1: the list has some elements
	if list.Size() != 3 {
		t.Error("case1 error: the list has some elements")
	}

	list = New(1)

	// case2: the list has one element
	if list.Size() != 1 {
		t.Error("case2 error: the list has one element")
	}

	list = New()

	// case3: the list has no element
	if list.Size() != 0 {
		t.Error("case3 error: the list has no element")
	}
}
func TestClear(t *testing.T) {
	list := New(1, 2, 3, 4)

	// case1: the list has some elements
	list.Clear()
	if list.size != 0 && list.first != nil && list.last != nil {
		t.Error("case1 error: the list has some elements")
	}

	list = New(1)

	// case2: the list has one element
	list.Clear()
	if list.size != 0 && list.first != nil && list.last != nil {
		t.Error("case2 error: the list has one element")
	}

	list = New()

	// case3: the list has no element
	list.Clear()
	if list.size != 0 && list.first != nil && list.last != nil {
		t.Error("case2 error: case3: the list has no element")
	}
}

func TestValues(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// case1: list has some elements
	result := []interface{}{1, 2, 3, 4, 5}
	values := list.Values()
	for i := range result {
		if result[i] != values[i] {
			t.Error("case1 error: list has some elements")
		}
	}

	list = New(1)

	// case2: list has one element
	result = []interface{}{1}
	values = list.Values()
	for i := range result {
		if result[i] != values[i] {
			t.Error("case2 error: list has one element")
		}
	}

	list = New()

	// case3: list has no element
	result = []interface{}{}
	values = list.Values()
	for i := range result {
		if result[i] != values[i] {
			t.Error("case3 error: list has no element")
		}
	}
}

// Serialization Interface

func TestToJSON(t *testing.T) {
	list := New(1, 2, 3, 4)

	// case1: list has some elements
	_, err := list.ToJSON()
	if err != nil {
		t.Error("case1 error: list has some elements")
	}

	list = New(1)

	// case2: list has one element
	_, err = list.ToJSON()
	if err != nil {
		t.Error("case2 error: list has one element")
	}

	list = New()

	// case3: list has no element
	_, err = list.ToJSON()
	if err != nil {
		t.Error("case3 error: list has no element")
	}
}

func TestFromJSON(t *testing.T) {
	list := New(1, 2, 3, 4)

	// case1: list has some elements
	json, _ := list.ToJSON()
	err := list.FromJSON(json)
	if err != nil {
		t.Error("case1 error: list has some elements")
	}

	list = New(1)

	// case2: list has one element
	json, _ = list.ToJSON()
	err = list.FromJSON(json)
	if err != nil {
		t.Error("case2 error: list has one element")
	}

	list = New()

	// case3: list has no element
	json, _ = list.ToJSON()
	err = list.FromJSON(json)
	if err != nil {
		t.Error("case3 error: list has no element")
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

func TestBegin(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)

	// case1: list has some elements
	iterator := list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 && iterator.element != nil {
		t.Error("case1 error: list has some elements")
	}

	list = New(1)

	// case2: list has one element
	iterator = list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 && iterator.element != nil {
		t.Error("case2 error: list has one element")
	}

	list = New()

	// case3: list has no element
	iterator = list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 && iterator.element != nil {
		t.Error("case3 error: list has no element")
	}
}

func TestNext(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)

	// case1: list has some elements
	iterator := list.Iterator()
	for iterator.Next() {
	}
	if iterator.index != 5 && iterator.element != nil {
		t.Error("case1 error: list has some elements")
	}

	list = New(1)

	// case2: list has one element
	iterator = list.Iterator()
	for iterator.Next() {
	}
	if iterator.index != 0 && iterator.element != nil {
		t.Error("case2 error: list has one element")
	}

	list = New()

	// case3: list has no element
	iterator = list.Iterator()
	for iterator.Next() {
	}
	if iterator.index != -1 && iterator.element != nil {
		t.Error("case3 error: list has no element")
	}
}

func TestValue(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)

	// case1: list has some elements
	iterator := list.Iterator()
	v := 1
	for iterator.Next() {
		if iterator.Value().(int) != v {
			t.Error("case1 error: list has some elements")
		}
		v++
	}

	list = New(1)

	// case2: list has one element
	iterator = list.Iterator()
	v = 1
	for iterator.Next() {
		if iterator.Value().(int) != v {
			t.Error("case2 error: list has one element")
		}
		v++
	}

	list = New()

	// case3: list has no element
	iterator = list.Iterator()
	iterator.Next()
	if iterator.Value() != nil {
		t.Error("case3 error: list has no element")
	}
}

func TestIndex(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6)

	// case1: list has some elements
	iterator := list.Iterator()
	i := 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case1 error: list has some elements")
		}
		i++
	}

	list = New(1)

	// case2: list has one element
	iterator = list.Iterator()
	i = 1
	for iterator.Next() {
		if iterator.Value().(int) != i {
			t.Error("case2 error: list has one element")
		}
		i++
	}

	list = New()

	// case3: list has no element
	iterator = list.Iterator()
	iterator.Next()
	if iterator.Index() != -1 {
		t.Error("case3 error: list has no element")
	}
}
