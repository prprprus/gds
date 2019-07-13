// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package singlylinkedlist

import (
	"testing"
)

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
	result := []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
		t.Error("case3 error: create a new list with some elements")
	}
}

// List Interface

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
	result := []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
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
	result := []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
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
	result = []interface{}{2, 1}
	if list.size != 2 || !verifyElements(result, list) {
		t.Error("case5 error: the list has one element, pre-append one element")
	}

	// case6: the list has one element, pre-append some elements
	list = New(1)
	list.PreAppend(2, 3, 4, 5)
	result = []interface{}{2, 3, 4, 5, 1}
	if list.size != 5 || !verifyElements(result, list) {
		t.Error("case6: the list has one element, pre-append some elements")
	}

	// case7: the list has some elements, pre-append nothing
	list = New(1, 2, 3, 4)
	list.PreAppend()
	result = []interface{}{1, 2, 3, 4}
	if list.size != 4 || !verifyElements(result, list) {
		t.Error("case7 error: the list has some elements, pre-append nothing")
	}

	// case8: the list has some elements, pre-append one element
	list = New(1, 2, 3, 4)
	list.PreAppend(5)
	result = []interface{}{5, 1, 2, 3, 4}
	if list.size != 5 || !verifyElements(result, list) {
		t.Error("case8 error: the list has some elements, pre-append one element")
	}

	// case9: the list has some elements, pre-append some elements
	list = New(1, 2, 3, 4)
	list.PreAppend(5, 6, 7, 8)
	result = []interface{}{5, 6, 7, 8, 1, 2, 3, 4}
	if list.size != 8 || !verifyElements(result, list) {
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
	if err != nil || list.first.value.(int) != 3 || list.first.next.next.value.(int) != 1 {
		t.Error("case5 error: the list has some elements, i unequal j")
	}

	// case6: the list has some elements, i equal j
	list = New(1, 2, 3, 4, 5)
	err = list.Swap(1, 1)
	if err != nil || list.first.next.value.(int) != 2 {
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
	if err != nil || list.size != 1 || list.first.value != 1 {
		t.Error("case4 error: the list has one element, insert nothing")
	}

	// case5: the list has one element, insert one element
	list = New(1)
	err = list.Insert(0, 2)
	if err != nil || list.size != 2 || list.first.value != 1 || list.first.next.value != 2 {
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
	if err != nil || list.first.value != 2 {
		t.Error("case2 error: the list has one element")
	}
	list = New(1)
	err = list.Set(1, 2)
	if err == nil || list.first.value == 2 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements, set first element
	list = New(1, 2, 3, 4)
	err = list.Set(0, -1)
	if err != nil || list.first.value != -1 {
		t.Error("case3 error: the list has some elements, set first element")
	}

	// case4: the list has some elements, set last element
	list = New(1, 2, 3, 4)
	err = list.Set(3, -1)
	if err != nil || list.last.value != -1 {
		t.Error("case4 error: the list has some elements, set last element")
	}

	// case5: the list has some elements, set medium element
	list = New(1, 2, 3, 4)
	err = list.Set(2, -1)
	if err != nil || list.first.next.next.value != -1 {
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

func TestReverse(t *testing.T) {
	// case1: the list has no element
	list := New()
	list.Reverse()
	if list.size != 0 || list.first != list.last {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	list.Reverse()
	if list.size != 1 || list.first != list.last || list.first.value != 1 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	list.Reverse()
	result := []interface{}{5, 4, 3, 2, 1}
	if list.size != 5 || !verifyElements(result, list) {
		t.Error("case3 error: the list has some elements")
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
	if list.size != 0 && list.first != nil && list.last != nil {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	list.Clear()
	if list.size != 0 && list.first != nil && list.last != nil {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	list.Clear()
	if list.size != 0 && list.first != nil && list.last != nil {
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
	if iterator.element != nil || iterator.index != -1 || list.first != list.last {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	if iterator.element != nil || iterator.index != -1 || list.first != list.last {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4)
	iterator = list.Iterator()
	if iterator.element != nil || iterator.index != -1 {
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
		if iterator.element.value != result[i] {
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
		if iterator.element.value != result[i] {
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
	if iterator.element != nil || iterator.index != -1 {
		t.Error("case1 error: the list has no element")
	}

	// case2: the list has one element
	list = New(1)
	iterator = list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.element != nil || iterator.index != -1 {
		t.Error("case2 error: the list has one element")
	}

	// case3: the list has some elements
	list = New(1, 2, 3, 4, 5)
	iterator = list.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.element != nil || iterator.index != -1 {
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

// Benchmark Test

// Append

func benchmarkAppend(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Append(i)
		}
	}
}

func BenchmarkAppend0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkAppend(b, list, size)
}

func BenchmarkAppend100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	b.StartTimer()
	benchmarkAppend(b, list, size)
}

func BenchmarkAppend1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	b.StartTimer()
	benchmarkAppend(b, list, size)
}

func BenchmarkAppend10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	b.StartTimer()
	benchmarkAppend(b, list, size)
}

func BenchmarkAppend100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	b.StartTimer()
	benchmarkAppend(b, list, size)
}

// Get

func benchmarkGet(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Get(i)
		}
	}
}

func BenchmarkGet0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkGet100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkGet1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkGet10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkGet100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

// Remove

func benchmarkRemove(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Remove(i)
		}
	}
}

func BenchmarkGetRemove0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkRemove100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkRemove1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkRemove10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkRemove100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

// Contains

func benchmarkContains(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Contains(i)
		}
	}
}

func BenchmarkGetContains0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkContains(b, list, size)
}

func BenchmarkContains100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkContains(b, list, size)
}

func BenchmarkContains1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkContains(b, list, size)
}

func BenchmarkContains10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkContains(b, list, size)
}

func BenchmarkContains100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkContains(b, list, size)
}

// Swap

func benchmarkSwap(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Swap(i, size-1)
		}
	}
}

func BenchmarkGetSwap0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkSwap(b, list, size)
}

func BenchmarkSwap100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSwap(b, list, size)
}

func BenchmarkSwap1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSwap(b, list, size)
}

func BenchmarkSwap10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSwap(b, list, size)
}

func BenchmarkSwap100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSwap(b, list, size)
}

// Insert

func benchmarkInsert(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Insert(i, i)
		}
	}
}

func BenchmarkGetInsert0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkInsert(b, list, size)
}

func BenchmarkInsert100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkInsert(b, list, size)
}

func BenchmarkInsert1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkInsert(b, list, size)
}

func BenchmarkInsert10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkInsert(b, list, size)
}

func BenchmarkInsert100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkInsert(b, list, size)
}

// Set

func benchmarkSet(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Set(i, i)
		}
	}
}

func BenchmarkGetSet0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkSet(b, list, size)
}

func BenchmarkSet100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSet(b, list, size)
}

func BenchmarkSet1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSet(b, list, size)
}

func BenchmarkSet10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSet(b, list, size)
}

func BenchmarkSet100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkSet(b, list, size)
}

// IndexOf

func benchmarkIndexOf(b *testing.B, list *List, size int) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.IndexOf(i)
		}
	}
}

func BenchmarkGetIndexOf0(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 0
	b.StartTimer()
	benchmarkIndexOf(b, list, size)
}

func BenchmarkIndexOf100(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkIndexOf(b, list, size)
}

func BenchmarkIndexOf1000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 1000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkIndexOf(b, list, size)
}

func BenchmarkIndexOf10000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 10000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkIndexOf(b, list, size)
}

func BenchmarkIndexOf100000(b *testing.B) {
	b.StopTimer()
	list := New()
	size := 100000
	for n := 0; n < size; n++ {
		list.Append(n)
	}
	b.StartTimer()
	benchmarkIndexOf(b, list, size)
}
