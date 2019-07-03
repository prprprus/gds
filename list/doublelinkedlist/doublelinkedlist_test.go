// test principle:
// 1. the list has no element
// 2. the list has one element
// 3. the list has some elements(handle first element, last element and medium element)

package doublelinkedlist

import "testing"

func verifyElements(result []interface{}, list *List) bool {
	flag := list.first
	for _, v := range result {
		if flag.value != v {
			return false
		}
		flag = flag.next
	}
	return true
}

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
	if !verifyElements(result, list) {
		t.Error("case3 error: create a new list with some elements")
	}
}

func TestAppend(t *testing.T) {
	// case1: the list has no element, append nothing
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
	if !verifyElements(result, list) {
		t.Error("case3 error: the list has no element, append some elements")
	}

	// case4: the list has one element, append nothing
	list = New(1)
	list.Append()
	if list.first != list.last || list.first.value != 1 || list.size != 1 {
		t.Error("case4 error: the list has one element, append an empty element")
	}

	// case5: the list has one element, append one element
	list = New(1)
	list.Append(2)
	if list.size != 2 {
		t.Error("case5 error: the list has one element, append one element")
	}
	result = []interface{}{1, 2}
	if !verifyElements(result, list) {
		t.Error("case5 error: the list has one element, append one element")
	}

	// case6: the list has one element, append some elements
	list = New(1)
	list.Append(2, 3, 4)
	if list.size != 4 {
		t.Error("case6 error: the list has one element, append some elements")
	}
	result = []interface{}{1, 2, 3, 4}
	if !verifyElements(result, list) {
		t.Error("case6 error: the list has one element, append some elements")
	}

	// case7: the list has some elements, append nothing
	list = New(1, 2, 3, 4)
	list.Append()
	if list.size != 4 {
		t.Error("case7 error: the list has some elements, append nothing")
	}

	// case8: the list has some elements, append one element
	list = New(1, 2, 3, 4)
	list.Append(5)
	if list.size != 5 {
		t.Error("case8 error: the list has some elements, append one element")
	}
	result = []interface{}{1, 2, 3, 4}
	if !verifyElements(result, list) {
		t.Error("case8 error: the list has some elements, append one element")
	}

	// case9: the list has some elements, append some elements
	list = New(1, 2, 3, 4)
	list.Append(5, 6, 7, 8)
	if list.size != 8 {
		t.Error("case9 error: the list has some elements, append some elements")
	}
	result = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	if !verifyElements(result, list) {
		t.Error("case9 error: the list has some elements, append some elements")
	}
}

func TestPreAppend(t *testing.T) {
	// case1: the list has no element, pre-append nothing
	list := New()
	list.PreAppend()
	if list.first != nil || list.last != nil || list.size != 0 {
		t.Error("case1 error: the list has no element, pre-append nothing")
	}

	// case2: the list has no element, pre-append one element
	list = New()
	list.PreAppend(1)
	if list.first != list.last || list.size != 1 || list.first.value != 1 {
		t.Error("case2 error: the list has no element, pre-append one element")
	}

	// case3: the list has no element, pre-append some elements
	list = New()
	list.PreAppend(1, 2, 3, 4)
	if list.size != 4 {
		t.Error("case3 error: the list has no element, pre-append some elements")
	}
	result := []interface{}{1, 2, 3, 4}
	if !verifyElements(result, list) {
		t.Error("case3 error: the list has no element, pre-append some elements")
	}

	// case4: the list has one element, pre-append nothing
	list = New(1)
	list.PreAppend()
	if list.first != list.last || list.size != 1 || list.first.value != 1 {
		t.Error("case4 error: the list has one element, pre-append nothing")
	}

	// case5: the list has one element, pre-append one element
	list = New(1)
	list.PreAppend(2)
	if list.size != 2 {
		t.Error("case5 error: the list has one element, pre-append one element")
	}
	result = []interface{}{2, 1}
	if !verifyElements(result, list) {
		t.Error("case5 error: the list has one element, pre-append one element")
	}

	// case6: the list has one element, pre-append some elements
	list = New(1)
	list.PreAppend(2, 3, 4, 5)
	if list.size != 5 {
		t.Error("case6: the list has one element, pre-append some elements")
	}
	result = []interface{}{2, 3, 4, 5, 1}
	if !verifyElements(result, list) {
		t.Error("case6: the list has one element, pre-append some elements")
	}

	// case7: the list has some elements, pre-append nothing
	list = New(1, 2, 3, 4)
	list.PreAppend()
	if list.size != 4 {
		t.Error("case7 error: the list has some elements, pre-append nothing")
	}
	result = []interface{}{1, 2, 3, 4}
	if !verifyElements(result, list) {
		t.Error("case7 error: the list has some elements, pre-append nothing")
	}

	// case8: the list has some elements, pre-append one element
	list = New(1, 2, 3, 4)
	list.PreAppend(5)
	if list.size != 5 {
		t.Error("case8 error: the list has some elements, pre-append one element")
	}
	result = []interface{}{5, 1, 2, 3, 4}
	if !verifyElements(result, list) {
		t.Error("case8 error: the list has some elements, pre-append one element")
	}

	// case9: the list has some elements, pre-append some elements
	list = New(1, 2, 3, 4)
	list.PreAppend(5, 6, 7, 8)
	if list.size != 8 {
		t.Error("case9 error: the list has some elements, pre-append some elements")
	}
	result = []interface{}{5, 6, 7, 8, 1, 2, 3, 4}
	if !verifyElements(result, list) {
		t.Error("case9 error: the list has some elements, pre-append some elements")
	}
}

func TestIndexInRange(t *testing.T) {
	// case1: the list has no element
	list := New()
	if list.indexInRange(0) {
		t.Error("case1 error: the list has no element")
	}
	if list.indexInRange(1) {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	if !list.indexInRange(0) {
		t.Error("case2 error: the list has one element")
	}
	if list.indexInRange(1) {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements, check with first element index
	list = New(1, 2, 3, 4, 5)
	if !list.indexInRange(0) {
		t.Error("case3 error: the list has some elements, check with first element index")
	}

	// case4: the list has some elements, check with last element index
	list = New(1, 2, 3, 4, 5)
	if !list.indexInRange(4) {
		t.Error("case4 error: the list has some elements, check with last element index")
	}

	// case5: the list has some elements, check with medium element index
	list = New(1, 2, 3, 4, 5)
	if !list.indexInRange(2) {
		t.Error("case4 error: the list has some elements, check with last element index")
	}
}

func TestGet(t *testing.T) {
	// case1: the list has no element
	list := New()
	_, err := list.Get(0)
	if err == nil {
		t.Error("case1 error: the list has no element")
	}
	_, err = list.Get(1)
	if err == nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	_, err = list.Get(0)
	if err != nil {
		t.Error("case2 error: the list has one element")
	}
	_, err = list.Get(1)
	if err == nil {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements, get first element
	list = New(1, 2, 3, 4)
	v, err := list.Get(0)
	if err != nil || v.(int) != 1 {
		t.Error("case3 error: the list has some elements, get first element")
	}

	// case4: the list has some elements, get last element
	list = New(1, 2, 3, 4)
	v, err = list.Get(3)
	if err != nil || v.(int) != 4 {
		t.Error("case4 error: the list has some elements, get last element")
	}

	// case5: the list has some elements, get medium element
	list = New(1, 2, 3, 4)
	v, err = list.Get(2)
	if err != nil || v.(int) != 3 {
		t.Error("case5 error: the list has some elements, get medium element")
	}
}

func TestRemove(t *testing.T) {
	// case1: the list has no element
	list := New()
	err := list.Remove(0)
	if err == nil {
		t.Error("case1 error: the list has no element")
	}
	err = list.Remove(1)
	if err == nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	err = list.Remove(0)
	if err != nil {
		t.Error("case2 error: the list has one element")
	}
	list = New(1)
	err = list.Remove(1)
	if err == nil {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements, remove first element
	list = New(1, 2, 3, 4, 5, 6)
	err = list.Remove(0)
	if err != nil {
		t.Error("case3 error: the list has some elements, remove first element")
	}

	// case4: the list has some elements, remove last element
	list = New(1, 2, 3, 4, 5, 6)
	err = list.Remove(5)
	if err != nil {
		t.Error("case4 error: the list has some elements, remove last element")
	}

	// case5: the list has some elements, remove medium element
	list = New(1, 2, 3, 4, 5, 6)
	err = list.Remove(3)
	if err != nil {
		t.Error("case5 error: the list has some elements, remove medium element")
	}
	err = list.Remove(2)
	if err != nil {
		t.Error("case5 error: the list has some elements, remove medium element")
	}
}

func TestContains(t *testing.T) {
	// case1: the list has no element, sub-list has no element
	list := New()
	if !list.Contains() {
		t.Error("case1 error: the list has no element, sub-list has no element")
	}

	// case2: the list has no element, sub-list has one element
	list = New()
	if list.Contains(1) {
		t.Error("case2 error: the list has no element, sub-list has one element")
	}

	// case3: the list has no element, sub-list has some elements
	list = New()
	if list.Contains(1, 2, 3) {
		t.Error("case1 error: the list has no element, sub-list has some elements")
	}

	// case4: the list has one element, sub-list has no element
	list = New(1)
	if !list.Contains() {
		t.Error("case4 error: the list has one element, sub-list has no element")
	}

	// case5: the list has one element, sub-list has one element and matching
	list = New(1)
	if !list.Contains(1) {
		t.Error("case5 error: the list has one element, sub-list has one element and matching")
	}

	// case6: the list has one element, sub-list has one element and not matching
	list = New(1)
	if list.Contains(2) {
		t.Error("case6 error: the list has one element, sub-list has one element and not matching")
	}

	// case7: the list has one element, sub-list has some elements
	list = New(1)
	if list.Contains(1, 2, 3) {
		t.Error("case7 error: the list has one element, sub-list has some element")
	}

	// case8: the list has some elements, sub-list has no element
	list = New(1, 2, 3, 4)
	if !list.Contains() {
		t.Error("case8 error: the list has some elements, sub-list has no element")
	}

	// case9: the list has some elements, sub-list has first element
	list = New(1, 2, 3, 4)
	if !list.Contains(1) {
		t.Error("case9 error: the list has some elements, sub-list has first element")
	}

	// case10: the list has some elements, sub-list has last element
	list = New(1, 2, 3, 4)
	if !list.Contains(4) {
		t.Error("case10 error: the list has some elements, sub-list has last element")
	}

	// case11: the list has some elements, sub-list has medium element
	list = New(1, 2, 3, 4)
	if !list.Contains(2, 3) {
		t.Error("case11 error: the list has some elements, sub-list has medium element")
	}
	if !list.Contains(1, 2, 3) {
		t.Error("case11 error: the list has some elements, sub-list has medium element")
	}
	if !list.Contains(2, 3, 4) {
		t.Error("case11 error: the list has some elements, sub-list has medium element")
	}
	if !list.Contains(1, 2, 3, 4) {
		t.Error("case11 error: the list has some elements, sub-list has medium element")
	}
}
