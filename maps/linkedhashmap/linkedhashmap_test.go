// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package linkedhashmap

import "testing"

// Map Interface

func TestPut(t *testing.T) {
	// case1: map has some elements
	m := New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if len(m.m) != 3 || m.m[1] != 1 || m.m[2] != 2 || m.m[3] != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New()
	m.Put(1, 1)
	if len(m.m) != 1 || m.m[1] != 1 {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New()
	if len(m.m) != 0 {
		t.Error("case3 error: map has one element")
	}
}

func TestGet(t *testing.T) {
	// case1: map has some elements
	m := New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	v1, _ := m.Get(1)
	v2, _ := m.Get(2)
	v3, _ := m.Get(3)
	if len(m.m) != 3 || v1.(int) != 1 || v2.(int) != 2 || v3.(int) != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New()
	m.Put(1, 1)
	v1, _ = m.Get(1)
	if len(m.m) != 1 || v1.(int) != 1 {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New()
	_, err := m.Get(1)
	if len(m.m) != 0 || err == nil {
		t.Error("case3 error: map has one element")
	}
}

func TestRemove(t *testing.T) {
	// case1: map has some elements
	m := New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Remove(2)
	m.Remove(1)
	_, err1 := m.Get(1)
	_, err2 := m.Get(2)
	if len(m.m) != 1 || err1 == nil || err2 == nil || m.m[3] != 3 {
		t.Error("case1 error: map has some elements")
	}

	// case2: map has one element
	m = New()
	m.Put(1, 1)
	m.Remove(1)
	_, err := m.Get(1)
	if len(m.m) != 0 || err == nil {
		t.Error("case2 error: map has one element")
	}

	// case3: map has no element
	m = New()
	m.Remove(1)
	if len(m.m) != 0 {
		t.Error("case3 error: map has one element")
	}
}

// Container Interface

func TestEmpty(t *testing.T) {
	// case1: the map has no element
	m := New()
	if !m.Empty() {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New()
	m.Put(1, 1)
	if m.Empty() {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if m.Empty() {
		t.Error("case3 error: the map has some elements")
	}
}

func TestSize(t *testing.T) {
	// case1: the map has no element
	m := New()
	if m.Size() != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New()
	m.Put(1, 1)
	if m.Size() != 1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	if m.Size() != 3 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestClear(t *testing.T) {
	// case1: the map has no element
	m := New()
	m.Clear()
	if len(m.m) != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the map has one element
	m = New()
	m.Put(1, 1)
	m.Clear()
	if len(m.m) != 0 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the map has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Clear()
	if len(m.m) != 0 {
		t.Error("case3 error: the map has some elements")
	}
}

func TestValues(t *testing.T) {
	// case1: the map has no element
	m := New()
	values := m.Values()
	if len(values) != 0 {
		t.Error("case1 error: the map has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	values = m.Values()
	if len(values) != 1 {
		t.Error("case2 error: the map has one element")
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	values = m.Values()
	if len(values) != 3 {
		t.Error("case3 error: the map has some elements")
	}
}

// Iterator Interface

func TestIterator(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	if iterator.m == nil || iterator.internalIterator == nil {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	if iterator.m == nil || iterator.internalIterator == nil {
		t.Error("case2 error: the m has one element")
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	if iterator.m == nil || iterator.internalIterator == nil {
		t.Error("case3 error: the m has some elements")
	}
}

func TestNext(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	if iterator.Next() {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Next() {
		if iterator.m[iterator.internalIterator.Value()] != result[i] {
			t.Error("case2 error: the m has one element")
		}
		i++
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	result = []interface{}{1, 2, 3, 4}
	i = 0
	for iterator.Next() {
		if iterator.m[iterator.internalIterator.Value()] != result[i] {
			t.Error("case3 error: the m has some elements")
		}
		i++
	}
}

func TestBegin(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case2 error: the m has one element")
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	for iterator.Next() {
	}
	iterator.Begin()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case3 error: the m has some elements")
	}
}

func TestPrev(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	if iterator.Prev() {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Prev() {
		if iterator.m[iterator.internalIterator.Value()] != result[i] {
			t.Error("case2 error: the m has one element")
		}
		i++
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	result = []interface{}{4, 3, 2, 1}
	i = 3
	for iterator.Prev() {
		if iterator.m[iterator.internalIterator.Value()] != result[i] {
			t.Error("case3 error: the m has some elements")
		}
		i++
	}
}

func TestEnd(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	for iterator.Prev() {
	}
	iterator.End()
	if iterator.internalIterator.Index() != 0 {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	for iterator.Prev() {
	}
	iterator.End()
	if iterator.internalIterator.Index() != 1 {
		t.Error("case2 error: the m has one element")
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	for iterator.Prev() {
	}
	iterator.End()
	if iterator.internalIterator.Index() != 4 {
		t.Error("case3 error: the m has some elements")
	}
}

func TestValue(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	iterator.Next()
	if iterator.Value() != nil {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	result := []interface{}{1}
	i := 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case2 error: the m has one element")
		}
		i++
	}

	// case3: the m has some elements
	m = New()
	iterator = m.Iterator()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	result = []interface{}{1, 2, 3, 4}
	i = 0
	for iterator.Next() {
		if iterator.Value() != result[i] {
			t.Error("case3 error: the m has some elements")
		}
		i++
	}
}

func TestKey(t *testing.T) {
	// case1: the m has no element
	m := New()
	iterator := m.Iterator()
	if iterator.internalIterator.Index() != -1 {
		t.Error("case1 error: the m has no element")
	}

	// case2: the m has one element
	m = New()
	m.Put(1, 1)
	iterator = m.Iterator()
	i := 0
	for iterator.Next() {
		if iterator.internalIterator.Index() != i {
			t.Error("case2 error: the m has one element")
		}
		i++
	}

	// case3: the m has some elements
	m = New()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(4, 4)
	iterator = m.Iterator()
	i = 0
	for iterator.Next() {
		if iterator.internalIterator.Index() != i {
			t.Error("case3 error: the m has some elements")
		}
		i++
	}
}
