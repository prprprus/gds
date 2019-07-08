// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

// Test Principle:
// 1. the list has no element.
// 2. the list has one element.
// 3. the list has some elements (handle first element, last element and medium element).

package arraylist

import "testing"

func verifyElements(result []interface{}, list *List) bool {
	for i, v := range result {
		if list.elements[i] != v {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {
	// case1: create a new list with no element
	list := New()
	if list.elements != nil || list.size != 0 {
		t.Error("case1 error: create a new list with no element")
	}

	// case2: create a new list with one element
	list = New(1)
	if list.elements == nil || list.size != 1 || list.elements[0] != 1 {
		t.Error("case2 error: create a new list with one element")
	}

	// case3: create a new list with some elements
	list = New(1, 2, 3, 4)
	result := []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
		t.Error("case3 error: create a new list with some elements")
	}
}

func TestGrowth(t *testing.T) {
	// case1: create a new list with no element
	list := New()
	if list.caps != 0 {
		t.Error("case1: create a new list with no element")
	}

	// case2: create a new list with one element
	list = New(1)
	if list.caps != 2 {
		t.Error("case2 error: create a new list with one element")
	}

	// case3: create a new list with some elements
	list = New(1, 2, 3, 4)
	if list.caps != 8 {
		t.Error("case3 error: create a new list with some elements")
	}
	list.Append(5, 6)
	if list.caps != 8 {
		t.Error("case3 error: create a new list with some elements")
	}
	list.Append(7, 8)
	if list.caps != 20 {
		t.Error("case3 error: create a new list with some elements")
	}
}

func TestShrink(t *testing.T) {
	// case1: create a new list with one element
	list := New(1)
	list.Remove(0)
	if list.caps != 0 {
		t.Error("case1: create a new list with no element")
	}

	// case2: create a new list with some elements
	list1 := New(1, 2, 3, 4, 5, 6, 7, 8)
	for i := 0; i < 7; i++ {
		list1.Remove(0)
	}
	if list1.caps != 1 {
		t.Error("case3 error: create a new list with some elements")
	}
}

// List Interface

func TestAppend(t *testing.T) {
	// case1: the list has no element, append nothing
	list := New()
	list.Append()
	if len(list.elements) != 0 || list.size != 0 {
		t.Error("case1 error: the list has no element, append an empty element")
	}

	// case2: the list has no element, append one element
	list = New()
	list.Append(1)
	if list.elements == nil || list.elements[0] != 1 || list.size != 1 {
		t.Error("case2 error: the list has no element, append one element")
	}

	// case3: the list has no element, append some elements
	list = New()
	list.Append(1, 2, 3, 4)
	result := []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
		t.Error("case3 error: the list has no element, append some elements")
	}

	// case4: the list has one element, append nothing
	list = New(1)
	list.Append()
	if list.elements == nil || list.elements[0] != 1 || list.size != 1 {
		t.Error("case4 error: the list has one element, append an empty element")
	}

	// case5: the list has one element, append one element
	list = New(1)
	list.Append(2)
	result = []interface{}{1, 2}
	if list.size != 2 || !verifyElements(result, list) {
		t.Error("case5 error: the list has one element, append one element")
	}

	// case6: the list has one element, append some elements
	list = New(1)
	list.Append(2, 3, 4)
	result = []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
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
	result = []interface{}{1, 2, 3, 4}
	if list.size != 5 || !verifyElements(result, list) {
		t.Error("case8 error: the list has some elements, append one element")
	}

	// case9: the list has some elements, append some elements
	list = New(1, 2, 3, 4)
	list.Append(5, 6, 7, 8)
	result = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	if list.size != 8 || !verifyElements(result, list) {
		t.Error("case9 error: the list has some elements, append some elements")
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

// func TestRemove(t *testing.T) {
// 	// case1: the list has no element
// 	list := New()
// 	err := list.Remove(0)
// 	if err == nil {
// 		t.Error("case1 error: the list has no element")
// 	}
// 	err = list.Remove(1)
// 	if err == nil {
// 		t.Error("case1 error: the list has no element")
// 	}

// 	// case2: the list has one element
// 	list = New(1)
// 	err = list.Remove(0)
// 	if err != nil {
// 		t.Error("case2 error: the list has one element")
// 	}
// 	list = New(1)
// 	err = list.Remove(1)
// 	if err == nil {
// 		t.Error("case2 error: the list has one element")
// 	}

// 	// case3: the list has some elements, remove first element
// 	list = New(1, 2, 3, 4, 5, 6)
// 	err = list.Remove(0)
// 	if err != nil {
// 		t.Error("case3 error: the list has some elements, remove first element")
// 	}

// 	// case4: the list has some elements, remove last element
// 	list = New(1, 2, 3, 4, 5, 6)
// 	err = list.Remove(5)
// 	if err != nil {
// 		t.Error("case4 error: the list has some elements, remove last element")
// 	}

// 	// case5: the list has some elements, remove medium element
// 	list = New(1, 2, 3, 4, 5, 6)
// 	err = list.Remove(3)
// 	if err != nil {
// 		t.Error("case5 error: the list has some elements, remove medium element")
// 	}
// 	err = list.Remove(2)
// 	if err != nil {
// 		t.Error("case5 error: the list has some elements, remove medium element")
// 	}
// }

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

func TestSwap(t *testing.T) {
	// case1: the list has no element, i unequal j
	list := New()
	err := list.Swap(0, 3)
	if err == nil {
		t.Error("case1 error: the list has no element, i unequal j")
	}

	// case2: the list has no element, i equal j
	list = New()
	err = list.Swap(0, 0)
	if err == nil {
		t.Error("case1 error: the list has no element, i unequal j")
	}

	// case3: the list has one element, i unequal j
	list = New(1)
	err = list.Swap(0, 2)
	if err == nil {
		t.Error("case2 error: the list has one element, i unequal j")
	}

	// case4: the list has one element, i equal j
	list = New(1)
	err = list.Swap(0, 0)
	if err != nil {
		t.Error("case4 error: the list has one element, i equal j")
	}
	err = list.Swap(1, 1)
	if err == nil {
		t.Error("case4 error: the list has one element, i equal j")
	}

	// case5: the list has some elements, i unequal j
	list = New(1, 2, 3, 4, 5)
	err = list.Swap(0, 2)
	if err != nil || list.elements[0].(int) != 3 || list.elements[2].(int) != 1 {
		t.Error("case5 error: the list has some elements, i unequal j")
	}

	// case6: the list has some elements, i equal j
	list = New(1, 2, 3, 4, 5)
	err = list.Swap(1, 1)
	if err != nil || list.elements[1].(int) != 2 {
		t.Error("case6 error: the list has some elements, i equal j")
	}
}

func TestInsert(t *testing.T) {
	// case1: the list has no element, insert nothing
	list := New()
	err := list.Insert(0)
	if err == nil {
		t.Error("case1 error: the list has no element, insert nothing")
	}

	// case2: the list has no element, insert one element
	list = New()
	err = list.Insert(0, 1)
	if err == nil {
		t.Error("case2 error: the list has no element, insert one element")
	}

	// case3: the list has no element, insert some elements
	list = New()
	err = list.Insert(0, 1, 2, 3, 4)
	if err == nil {
		t.Error("case3 error: the list has no element, insert some elements")
	}

	// case4: the list has one element, insert nothing
	list = New(1)
	err = list.Insert(0)
	if err != nil || list.size != 1 || list.elements[0] != 1 {
		t.Error("case4 error: the list has one element, insert nothing")
	}

	// case5: the list has one element, insert one element
	list = New(1)
	err = list.Insert(0, 2)
	if err != nil || list.size != 2 || list.elements[0] != 1 || list.elements[1] != 2 {
		t.Error("case5 error: the list has one element, insert one element")
	}
	list = New(1)
	err = list.Insert(1, 2)
	if err == nil {
		t.Error("case5 error: the list has one element, insert one element")
	}

	// case6: the list has one element, insert some elements
	list = New(1)
	err = list.Insert(0, 2, 3, 4, 5, 6)
	result := []interface{}{1, 2, 3, 4, 5, 6}
	if err != nil || list.size != 6 || !verifyElements(result, list) {
		t.Error("case6 error: the list has one element, insert some elements")
	}

	// case7: the list has some elements, insert nothing
	list = New(1, 2, 3, 4)
	err = list.Insert(0)
	result = []interface{}{1, 2, 3, 4}
	if err != nil || list.size != 4 || !verifyElements(result, list) {
		t.Error("case7 error: the list has some elements, insert nothing")
	}

	// case8: the list has some elements, insert one element after the first element
	list = New(1, 2, 3, 4)
	err = list.Insert(0, -1)
	result = []interface{}{1, -1, 2, 3, 4}
	if err != nil || list.size != 5 || !verifyElements(result, list) {
		t.Error("case8 error: the list has some elements, insert one element after the first element")
	}

	// case9: the list has some elements, insert one element after the last element
	list = New(1, 2, 3, 4)
	err = list.Insert(3, -1)
	result = []interface{}{1, 2, 3, 4, -1}
	if err != nil || list.size != 5 || !verifyElements(result, list) {
		t.Error("case9 error: the list has some elements, insert one element after the last element")
	}

	// case10: the list has some elements, insert one element after the medium element
	list = New(1, 2, 3, 4)
	err = list.Insert(2, -1)
	result = []interface{}{1, 2, 3, -1, 4}
	if err != nil || list.size != 5 || !verifyElements(result, list) {
		t.Error("case10 error: the list has some elements, insert one element after the medium element")
	}

	// case11: the list has some elements, insert some elements after the first element
	list = New(1, 2, 3, 4)
	err = list.Insert(0, -1, -2, -3, -4)
	result = []interface{}{1, -1, -2, -3, -4, 2, 3, 4}
	if err != nil || list.size != 8 || !verifyElements(result, list) {
		t.Error("case11 error: the list has some elements, insert some elements after the first element")
	}

	// case12: the list has some elements, insert some elements after the last element
	list = New(1, 2, 3, 4)
	err = list.Insert(3, -1, -2, -3, -4)
	result = []interface{}{1, 2, 3, 4, -1, -2, -3, -4}
	if err != nil || list.size != 8 || !verifyElements(result, list) {
		t.Error("case12 error: the list has some elements, insert some elements after the last element")
	}

	// case13: the list has some elements, insert some elements after the medium element
	list = New(1, 2, 3, 4)
	err = list.Insert(1, -1, -2, -3, -4)
	result = []interface{}{1, 2, -1, -2, -3, -4, 3, 4}
	if err != nil || list.size != 8 || !verifyElements(result, list) {
		t.Error("case13 error: the list has some elements, insert some elements after the medium element")
	}
}

func TestSet(t *testing.T) {
	// case1: the list has no element
	list := New()
	err := list.Set(0, 1)
	if err == nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	err = list.Set(0, 2)
	if err != nil || list.elements[0] != 2 {
		t.Error("case2 error: the list has one element")
	}
	list = New(1)
	err = list.Set(1, 2)
	if err == nil || list.elements[0] == 2 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements, set first element
	list = New(1, 2, 3, 4)
	err = list.Set(0, -1)
	if err != nil || list.elements[0] != -1 {
		t.Error("case3 error: the list has some elements, set first element")
	}

	// case4: the list has some elements, set last element
	list = New(1, 2, 3, 4)
	err = list.Set(3, -1)
	if err != nil || list.elements[3] != -1 {
		t.Error("case4 error: the list has some elements, set last element")
	}

	// case5: the list has some elements, set medium element
	list = New(1, 2, 3, 4)
	err = list.Set(2, -1)
	if err != nil || list.elements[2] != -1 {
		t.Error("case5 error: the list has some elements, set medium element")
	}
}

func TestIndexOf(t *testing.T) {
	// case1: the list has no element
	list := New()
	i, err := list.IndexOf(1)
	if i != -1 || err == nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element and matching
	list = New(1)
	i, err = list.IndexOf(1)
	if i != 0 || err != nil {
		t.Error("case2 error: the list has one element and matching")
	}

	// case3: the list has one element and not matching
	list = New(1)
	i, err = list.IndexOf(2)
	if i != -1 || err == nil {
		t.Error("case3 error: the list has one element and not matching")
	}

	// case4: the list has some elements, matching first element
	list = New(1, 2, 3, 4)
	i, err = list.IndexOf(1)
	if i != 0 || err != nil {
		t.Error("case4 error: the list has some elements, matching first element")
	}

	// case5: the list has some elements, matching last element
	list = New(1, 2, 3, 4)
	i, err = list.IndexOf(4)
	if i != 3 || err != nil {
		t.Error("case5 error: the list has some elements, matching last element")
	}

	// case6: the list has some elements, matching medium element
	list = New(1, 2, 3, 4)
	i, err = list.IndexOf(3)
	if i != 2 || err != nil {
		t.Error("case6 error: the list has some elements, matching medium element")
	}

	// case7: the list has some elements and not matching
	list = New(1, 2, 3, 4)
	i, err = list.IndexOf(5)
	if i != -1 || err == nil {
		t.Error("case7 error: the list has some elements and not matching")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: the list has no element
	list := New()
	if !list.Empty() {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	if list.Empty() {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	if list.Empty() {
		t.Error("case3 error: the list has some elements")
	}
}

func TestSize(t *testing.T) {
	// case1: the list has no element
	list := New()
	if list.Size() != 0 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	if list.Size() != 1 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	if list.Size() != 5 {
		t.Error("case3 error: the list has some elements")
	}
}

func TestClear(t *testing.T) {
	// case1: the list has no element
	list := New()
	list.Clear()
	if list.size != 0 && list.elements != nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	list.Clear()
	if list.size != 0 && list.elements != nil {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	list.Clear()
	if list.size != 0 && list.elements != nil {
		t.Error("case3 error: the list has some elements")
	}
}

func TestValues(t *testing.T) {
	// case1: the list has no element
	list := New()
	values := list.Values()
	if len(values) != 0 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	values = list.Values()
	if !verifyElements(values, list) {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	values = list.Values()
	if !verifyElements(values, list) {
		t.Error("case3 error: the list has some elements")
	}
}

// Iterator Interface

func TestIterator(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	if iterator.index != -1 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	if iterator.index != -1 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4)
	iterator = list.Iterator()
	if iterator.index != -1 {
		t.Error("case3 error: the list has some elements")
	}
}

func TestNext(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	if iterator.Next() {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Next() {
		if iterator.list.elements[i] != result[i] {
			t.Error("case2 error: the list has one element")
		}
		i++
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4)
	iterator = list.Iterator()
	result = []interface{}{1, 2, 3, 4}
	i = 0
	for iterator.Next() {
		if iterator.list.elements[i] != result[i] {
			t.Error("case3 error: the list has some elements")
		}
		i++
	}
}

func TestBegin(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	iterator = list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.index != -1 {
		t.Error("case3 error: the list has some elements")
	}
}

func TestPrev(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	if iterator.Prev() {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Prev() {
		if iterator.list.elements[i] != result[i] {
			t.Error("case2 error: the list has one element")
		}
		i++
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4)
	iterator = list.Iterator()
	result = []interface{}{4, 3, 2, 1}
	i = 3
	for iterator.Prev() {
		if iterator.list.elements[i] != result[i] {
			t.Error("case3 error: the list has some elements")
		}
		i++
	}
}

func TestEnd(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	for iterator.Prev() {
	}
	iterator.End()
	if iterator.index != 0 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	for iterator.Prev() {
	}
	iterator.End()
	if iterator.index != 1 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	iterator = list.Iterator()
	for iterator.Prev() {
	}
	iterator.End()
	if iterator.index != 5 {
		t.Error("case3 error: the list has some elements")
	}
}

func TestValue(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	iterator.Next()
	if iterator.Value() != nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case2 error: the list has one element")
		}
		i++
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4)
	iterator = list.Iterator()
	result = []interface{}{1, 2, 3, 4}
	i = 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case3 error: the list has some elements")
		}
		i++
	}
}

func TestIndex(t *testing.T) {
	// case1: the list has no element
	list := New()
	iterator := list.Iterator()
	if iterator.index != -1 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	i := 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case2 error: the list has one element")
		}
		i++
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4)
	iterator = list.Iterator()
	i = 0
	for iterator.Next() {
		if iterator.Index() != i {
			t.Error("case3 error: the list has some elements")
		}
		i++
	}
}
